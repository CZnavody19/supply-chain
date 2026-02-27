import { useEffect, useState } from 'react';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import { getOrders, getSupplyPath, getCostBreakdown, getOptimalRoute, getLocations, getCompanies } from '../api';
import StatusBadge from '../components/StatusBadge';
import { formatCurrency, formatNumber } from '../utils/formatters';

export default function OrdersPage() {
  const [orders, setOrders] = useState([]);
  const [path, setPath] = useState(null);
  const [costs, setCosts] = useState(null);
  
  const [statusFilter, setStatusFilter] = useState('');
  const [searchFilter, setSearchFilter] = useState('');
  
  const [routeInput, setRouteInput] = useState({ from: 'loc-001', to: 'loc-002' });
  const [routeError, setRouteError] = useState('');
  const [optimalRoute, setOptimalRoute] = useState(null);
  
  // Store a dictionary of id -> coordinates
  const [coordDictionary, setCoordDictionary] = useState({});

  useEffect(() => { 
    getOrders().then(setOrders).catch(console.error); 
    
    // Fetch all locations and companies to map their coordinates
    Promise.all([getLocations(), getCompanies()])
      .then(([locs, comps]) => {
        const dict = {};
        locs.forEach(l => dict[l.id] = l.coordinates || { lat: l.lat, lng: l.lng });
        comps.forEach(c => dict[c.id] = c.coordinates || { lat: c.lat, lng: c.lng });
        setCoordDictionary(dict);
      })
      .catch(console.error);
  }, []);

  const viewDetails = (id) => {
    getSupplyPath(id).then(setPath).catch(console.error);
    getCostBreakdown(id).then(setCosts).catch(console.error);
  };

  const calcRoute = () => {
    setRouteError('');
    setOptimalRoute(null);
    getOptimalRoute(routeInput.from, routeInput.to)
    .then(setOptimalRoute)
    .catch(err => {
      if (err.response && err.response.data) {
        setRouteError(typeof err.response.data === 'string' ? err.response.data : 'Route not found.');
      } else {
        setRouteError('An error occurred while calculating the route.');
      }
    });
  };

  const filteredOrders = orders.filter(o => {
    const matchesStatus = statusFilter ? o.status === statusFilter : true;
    const matchesSearch = o.id.toLowerCase().includes(searchFilter.toLowerCase());
    return matchesStatus && matchesSearch;
  });

  // Look up coordinates from the dictionary using the IDs in the supply path
  const mapMarkers = [];
  path?.path?.forEach(stage => {
    const targetId = stage.location?.id || stage.from || stage.company?.id;
    const name = stage.location?.name || stage.company?.name || stage.name;
    
    if (targetId && coordDictionary[targetId]) {
      const coords = coordDictionary[targetId];
      if (coords.lat && coords.lng) {
        mapMarkers.push({ position: [coords.lat, coords.lng], name });
      }
    }
  });

  const mapCenter = mapMarkers.length > 0 ? mapMarkers[0].position : [25.033, 121.565];

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
      <div className="grid-2">
        <div className="card">
          <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '1rem', alignItems: 'center' }}>
            <h2 style={{ margin: 0 }}>Orders</h2>
            <div style={{ display: 'flex', gap: '0.5rem' }}>
              <select onChange={e => setStatusFilter(e.target.value)} style={{ padding: '0.4rem' }}>
                <option value="">All Statuses</option>
                <option value="pending">Pending</option>
                <option value="in_transit">In Transit</option>
                <option value="delivered">Delivered</option>
                <option value="delayed">Delayed</option>
              </select>
              <input 
                type="text" 
                placeholder="Search Order ID..." 
                onChange={e => setSearchFilter(e.target.value)}
                style={{ padding: '0.4rem', width: '220px' }}
              />
            </div>
          </div>
          <table>
            <thead><tr><th>ID</th><th>Status</th><th>Quantity</th><th>Action</th></tr></thead>
            <tbody>
              {filteredOrders.map(o => (
                <tr key={o.id}>
                  <td>{o.id}</td><td><StatusBadge status={o.status} /></td><td>{o.quantity}</td>
                  <td><button className="btn" onClick={() => viewDetails(o.id)}>View Details</button></td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        <div className="card">
          <h3>Order Details & Logistics</h3>
          {path && costs ? (
            <div>
              <div style={{ background: '#f3f4f6', padding: '1rem', borderRadius: '4px', marginBottom: '1rem' }}>
                <strong>Cost Breakdown:</strong>
                <div>Material: {formatCurrency(costs.materialCost)} | Manufacturing: {formatCurrency(costs.manufacturingCost)} | Logistics: {formatCurrency(costs.logisticsCost)}</div>
                <div style={{ fontSize: '1.2rem', fontWeight: 'bold', marginTop: '0.5rem' }}>Total: {formatCurrency(costs.totalCost)}</div>
              </div>
              
              <p>Total Duration: {path.totalDuration}</p>
              <ul>{path.path.map((p, i) => <li key={i}><strong>{p.name}</strong> - {p.status}</li>)}</ul>
              <div style={{ height: '300px', marginTop: '1rem', zIndex: 0 }}>
                <MapContainer key={path.orderId} center={mapCenter} zoom={3} style={{ height: '100%', width: '100%' }}>
                  <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
                  {mapMarkers.map((marker, idx) => (
                    <Marker key={idx} position={marker.position}>
                      <Popup>{marker.name}</Popup>
                    </Marker>
                  ))}
                </MapContainer>
              </div>
            </div>
          ) : <p>Select an order to view logistics.</p>}
        </div>
      </div>

      <div className="card">
        <h3>Optimal Route Calculator</h3>
        <div style={{ display: 'flex', gap: '1rem', marginBottom: '1rem' }}>
          <input 
            value={routeInput.from} 
            onChange={e => setRouteInput({...routeInput, from: e.target.value})} 
            placeholder="From Location ID" 
            style={{ padding: '0.5rem' }} 
          />
          <input 
            value={routeInput.to} 
            onChange={e => setRouteInput({...routeInput, to: e.target.value})} 
            placeholder="To Location ID" 
            style={{ padding: '0.5rem' }} 
          />
          <button className="btn" onClick={calcRoute}>Calculate Route</button>
        </div>

        {routeError && (
          <div style={{ padding: '1rem', background: '#fee2e2', color: '#991b1b', borderRadius: '4px', marginBottom: '1rem' }}>
            {routeError}
          </div>
        )}

        {optimalRoute && (
          <div style={{ padding: '1rem', background: '#eef2ff', borderRadius: '4px' }}>
            <p>
              <strong>Distance:</strong> {formatNumber(optimalRoute.totalDistance)} kilometers | 
              <strong>Time:</strong> {formatNumber(optimalRoute.totalTime)} hours | 
              <strong>Cost:</strong> {formatCurrency(optimalRoute.totalCost)}
            </p>
          </div>
        )}
      </div>
    </div>
  );
}