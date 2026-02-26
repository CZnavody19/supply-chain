import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { getProducts, getProductBom } from '../api';
import StatusBadge from '../components/StatusBadge';

export default function ProductsPage() {
  const [products, setProducts] = useState([]);
  const [search, setSearch] = useState('');
  const [expandedId, setExpandedId] = useState(null);
  const [bomData, setBomData] = useState({});

  useEffect(() => { getProducts().then(setProducts).catch(console.error); }, []);

  const toggleBom = async (id) => {
    if (expandedId === id) {
      setExpandedId(null);
      return;
    }
    if (!bomData[id]) {
      const bom = await getProductBom(id);
      setBomData(prev => ({ ...prev, [id]: bom }));
    }
    setExpandedId(id);
  };

  const filtered = products.filter(p => p.name.toLowerCase().includes(search.toLowerCase()));

  return (
    <div className="card">
      <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '1rem' }}>
        <h2>Products Catalog</h2>
        <input 
          type="text" 
          placeholder="Search products..." 
          value={search} 
          onChange={(e) => setSearch(e.target.value)} 
          style={{ padding: '0.5rem', borderRadius: '4px', border: '1px solid #ccc' }}
        />
      </div>
      <table>
        <thead>
          <tr><th>Name</th><th>SKU</th><th>Price</th><th>Status</th><th>Actions</th></tr>
        </thead>
        <tbody>
          {filtered.map(p => (
            <div style={{ display: 'contents' }} key={p.id}>
              <tr>
                <td>{p.name}</td><td>{p.sku}</td><td>${p.price}</td>
                <td><StatusBadge status={p.status} /></td>
                <td style={{ display: 'flex', gap: '0.5rem' }}>
                  <button className="btn" onClick={() => toggleBom(p.id)}>
                    {expandedId === p.id ? 'Hide BOM' : 'Show BOM'}
                  </button>
                  <Link to={`/products/${p.id}`} className="btn" style={{textDecoration:'none', background:'#4b5563'}}>Details</Link>
                </td>
              </tr>
              {expandedId === p.id && bomData[p.id] && (
                <tr style={{ background: '#f8fafc' }}>
                  <td colSpan="5" style={{ padding: '1rem' }}>
                    <strong>BOM Structure:</strong>
                    <ul style={{ margin: '0.5rem 0 0 1rem' }}>
                      {bomData[p.id].map((item, i) => (
                        <li key={i}>{item.component?.name || 'Unknown'} (Qty: {item.quantity})</li>
                      ))}
                    </ul>
                  </td>
                </tr>
              )}
            </div>
          ))}
        </tbody>
      </table>
    </div>
  );
}