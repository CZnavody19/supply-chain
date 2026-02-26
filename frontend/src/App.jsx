import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import ProductsPage from './pages/ProductsPage';
import ProductDetailPage from './pages/ProductDetailPage';
import SuppliersPage from './pages/SuppliersPage';
import OrdersPage from './pages/OrdersPage';
import HealthDashboard from './pages/HealthDashboard';
import RiskScenarioPage from './pages/RiskScenarioPage';

export default function App() {
  return (
    <BrowserRouter>
      <div className="app-container">
        <nav className="sidebar">
          <h2 style={{ padding: '0 1rem', color: '#fff' }}>SCM Network</h2>
          <Link to="/analytics/health">Dashboard Health</Link>
          <Link to="/products">Products & BOM</Link>
          <Link to="/suppliers">Suppliers & Risk</Link>
          <Link to="/orders">Orders & Logistics</Link>
          <Link to="/analytics/scenarios">Risk Scenarios</Link>
        </nav>
        <main className="main-content">
          <Routes>
            <Route path="/" element={<HealthDashboard />} />
            <Route path="/analytics/health" element={<HealthDashboard />} />
            <Route path="/products" element={<ProductsPage />} />
            <Route path="/products/:id" element={<ProductDetailPage />} />
            <Route path="/suppliers" element={<SuppliersPage />} />
            <Route path="/orders" element={<OrdersPage />} />
            <Route path="/analytics/scenarios" element={<RiskScenarioPage />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  );
}