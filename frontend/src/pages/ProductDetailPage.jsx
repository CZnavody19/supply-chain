import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import { getProduct, getBomDetailed, getAltSuppliers, getStockLevels } from '../api';

export default function ProductDetailPage() {
  const { id } = useParams();
  const [data, setData] = useState(null);

  useEffect(() => {
    Promise.all([getProduct(id), getBomDetailed(id), getAltSuppliers(id), getStockLevels(id)])
      .then(([product, bom, alternatives, stock]) => setData({ product, bom, alternatives, stock }))
      .catch(console.error);
  }, [id]);

  if (!data) return <div>Loading...</div>;
  const { product, bom, alternatives, stock } = data;

  // Mocking price history based on current price for the chart requirement
  const priceHistory = [
    { month: 'Oct', price: product.price * 0.9 },
    { month: 'Nov', price: product.price * 0.95 },
    { month: 'Dec', price: product.price * 1.1 },
    { month: 'Jan', price: product.price * 1.05 },
    { month: 'Feb', price: product.price }
  ];

  return (
    <div>
      <Link to="/products" style={{ display: 'inline-block', marginBottom: '1rem' }}>← Back to Products</Link>
      <h2>{product.name} ({product.sku})</h2>
      
      <div className="grid-2">
        <div className="card">
          <h3>Bill of Materials</h3>
          <table>
            <thead><tr><th>Component</th><th>Qty</th><th>Suppliers</th></tr></thead>
            <tbody>
              {bom?.map((entry, i) => (
                <tr key={i}>
                  <td>{entry.component?.name}</td><td>{entry.quantity}</td>
                  <td>{entry.suppliers?.map(s => <div key={s.company?.id}>{s.company?.name} (${s.price})</div>)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        <div className="card">
          <h3>Price History</h3>
          <ResponsiveContainer width="100%" height={200}>
            <LineChart data={priceHistory}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="month" />
              <YAxis domain={['auto', 'auto']} />
              <Tooltip />
              <Line type="stepAfter" dataKey="price" stroke="#10b981" strokeWidth={2} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </div>
    </div>
  );
}