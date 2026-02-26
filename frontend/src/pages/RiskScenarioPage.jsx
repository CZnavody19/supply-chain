import { useEffect, useState } from 'react';
import { getCompanies, getImpact } from '../api';
import { formatCurrency } from '../utils/formatters';

export default function RiskScenarioPage() {
  const [suppliers, setSuppliers] = useState([]);
  const [impact, setImpact] = useState(null);

  useEffect(() => { 
    getCompanies()
      .then(res => setSuppliers(res.filter(c => c.type === 'supplier')))
      .catch(console.error); 
  }, []);

  const runScenario = (e) => {
    getImpact(e.target.value).then(setImpact).catch(console.error);
  };

  return (
    <div className="card">
      <h2>Impact Analysis Simulator</h2>
      <select onChange={runScenario} defaultValue="" style={{ padding: '0.5rem', marginBottom: '1rem', width: '300px' }}>
        <option value="" disabled>Select a supplier to fail...</option>
        {suppliers.map(s => <option key={s.id} value={s.id}>{s.name}</option>)}
      </select>

      {impact && impact.impact && (
        <div style={{ borderTop: '1px solid #e4e4e7', paddingTop: '1rem' }}>
          <h3>Impact of {impact.supplierName} Failure</h3>
          <p>
            Estimated Cost: <strong>{formatCurrency(impact.impact.estimatedCost)}</strong> | 
            Affected Revenue: <strong>{formatCurrency(impact.impact.affectedRevenue)}</strong>
          </p>
          
          <h4>Affected Products</h4>
          {impact.impact.affectedProducts?.length > 0 ? (
            <ul>
              {impact.impact.affectedProducts.map(p => (
                <li key={p.productId}>{p.productName} - Delayed by {p.delayDays} days</li>
              ))}
            </ul>
          ) : (
            <p>No products affected by this supplier.</p>
          )}

          <h4>Mitigation Strategies</h4>
          {impact.impact.mitigation?.length > 0 ? (
            <ul>
              {impact.impact.mitigation.map((m, i) => <li key={i}>{m}</li>)}
            </ul>
          ) : (
            <p>No immediate mitigation strategies available.</p>
          )}
        </div>
      )}
    </div>
  );
}