import { useEffect, useState } from 'react';
import { getCompanies, getRiskAssessment } from '../api';
import { formatPercent, formatNumber } from '../utils/formatters';

export default function SuppliersPage() {
  const [suppliers, setSuppliers] = useState([]);
  const [risk, setRisk] = useState(null);
  
  // New filter states
  const [countryFilter, setCountryFilter] = useState('');
  const [minReliability, setMinReliability] = useState(0);

  useEffect(() => {
    getCompanies()
      .then(res => setSuppliers(res.filter(c => c.type === 'supplier')))
      .catch(console.error);
  }, []);

  const filteredSuppliers = suppliers.filter(s => {
    const matchesCountry = s.country.toLowerCase().includes(countryFilter.toLowerCase());
    const matchesRel = s.reliability >= minReliability;
    return matchesCountry && matchesRel;
  });

  return (
    <div className="grid-2">
      <div className="card">
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1rem' }}>
          <h2>Suppliers</h2>
          <div style={{ display: 'flex', gap: '0.5rem' }}>
            <input placeholder="Filter country..." onChange={e => setCountryFilter(e.target.value)} style={{ padding: '0.4rem' }} />
            <input type="number" step="0.1" min="0" max="1" placeholder="Min Rating (0-1)" onChange={e => setMinReliability(Number(e.target.value) || 0)} style={{ width: '120px', padding: '0.4rem' }} />
          </div>
        </div>
        <table>
          <thead><tr><th>Name</th><th>Country</th><th>Reliability</th><th>Action</th></tr></thead>
          <tbody>
            {filteredSuppliers.map(s => (
              <tr key={s.id}>
                <td>{s.name}</td><td>{s.country}</td>
                <td>{formatPercent(s.reliability)}</td>
                <td><button className="btn" onClick={() => getRiskAssessment(s.id).then(setRisk)}>Analyze Risk</button></td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      
      {risk && (
        <div className="card" style={{ borderTop: '4px solid #3b82f6' }}>
          <h3>Risk Assessment: {risk.company}</h3>
          
          <div style={{ display: 'flex', alignItems: 'baseline', gap: '1rem', marginBottom: '1rem' }}>
            <h1 style={{ color: risk.riskScore * 100 < 30 ? '#166534' : '#991b1b', margin: 0 }}>
              Risk Score: {formatNumber(risk.riskScore * 100, 0)}/100
            </h1>
            <span style={{ fontSize: '0.9rem', color: '#6b7280' }}>(Lower is better)</span>
          </div>

          <ul>
            <li>Reliability: {formatPercent(risk.factors.reliabilityScore)}</li>
            <li>On Time Delivery: {formatPercent(risk.factors.onTimeDeliveryRate)}</li>
            <li>Financial Stability: {formatPercent(risk.factors.financialStability)}</li>
            <li>Geopolitical Risk: {formatPercent(risk.factors.geopoliticalRisk)}</li>
          </ul>
          
          <h4 style={{ marginTop: '1.5rem' }}>Recommendations</h4>
          <ul>{risk.recommendations.map((r, i) => <li key={i}>{r}</li>)}</ul>
        </div>
      )}
    </div>
  );
}