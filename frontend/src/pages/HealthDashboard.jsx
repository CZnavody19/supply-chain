import { useEffect, useState } from 'react';
import { getHealth, getForecastDelays } from '../api';
import StatusBadge from '../components/StatusBadge';
import { formatPercent, formatNumber } from '../utils/formatters';

export default function HealthDashboard() {
  const [health, setHealth] = useState(null);
  const [forecasts, setForecasts] = useState([]);

  useEffect(() => {
    getHealth().then(setHealth).catch(console.error);
    getForecastDelays(3).then(setForecasts).catch(console.error);
  }, []);

  if (!health) return <div style={{ padding: '2rem' }}>Loading dashboard...</div>;

  return (
    <div>
      <h2>Supply Chain Health</h2>
      <div className="grid-3" style={{ marginBottom: '1.5rem' }}>
        <div className="card" style={{ borderLeft: '4px solid #3b82f6', marginBottom: 0 }}>
          <h4 style={{ margin: 0, color: '#6b7280' }}>On-Time Delivery (Time)</h4>
          <h1 style={{ margin: '0.5rem 0 0' }}>92.4%</h1>
        </div>
        <div className="card" style={{ borderLeft: '4px solid #10b981', marginBottom: 0 }}>
          <h4 style={{ margin: 0, color: '#6b7280' }}>Avg Cost Variance (Cost)</h4>
          <h1 style={{ margin: '0.5rem 0 0' }}>+1.2%</h1>
        </div>
        <div className="card" style={{ borderLeft: '4px solid #8b5cf6', marginBottom: 0 }}>
          <h4 style={{ margin: 0, color: '#6b7280' }}>Defect Rate (Quality)</h4>
          <h1 style={{ margin: '0.5rem 0 0' }}>0.8%</h1>
        </div>
      </div>
      <div className="grid-3">
        <div className="card">
          <h3>Critical Components</h3>
          <ul>{health.criticalComponents?.map(c => <li key={c.componentId}>{c.componentName}</li>)}</ul>
        </div>
        <div className="card">
          <h3>Bottlenecks</h3>
          <ul>
            {health.bottlenecks?.map(b => (
              <li key={b.locationId}>{b.locationName} (Utilization: {formatPercent(b.utilization)})</li>
            ))}
          </ul>
        </div>
        <div className="card">
          <h3>High Risk Suppliers</h3>
          <ul>
            {health.highRiskSuppliers?.map(s => (
              <li key={s.companyId}>{s.companyName} (Reliability: {formatPercent(s.reliability)})</li>
            ))}
          </ul>
        </div>
      </div>

      <div className="grid-2">
        <div className="card">
          <h3>Recommendations</h3>
          <ul>{health.recommendations?.map((r, i) => <li key={i}>{r}</li>)}</ul>
        </div>

        <div className="card">
          <h3>3-Month Delay Forecast</h3>
          <table>
            <thead><tr><th>Product</th><th>Average Delay</th><th>Risk Level</th></tr></thead>
            <tbody>
              {forecasts?.map((f, i) => (
                <tr key={i}>
                  <td>{f.productName}</td>
                  <td>{formatNumber(f.avgDelayDays)} days</td>
                  <td><StatusBadge status={f.riskLevel === 'high' ? 'delayed' : 'active'} /></td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}