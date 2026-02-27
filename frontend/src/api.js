import axios from 'axios';

export const api = axios.create({ baseURL: 'http://localhost:8080/api' });

export const getProducts = () => api.get('/products').then(res => res.data);
export const getProduct = (id) => api.get(`/products?id=${id}`).then(res => res.data);
export const getProductBom = (id) => api.get(`/products/bom?id=${id}`).then(res => res.data);
export const getBomDetailed = (id) => api.get(`/products/bom/detailed?id=${id}`).then(res => res.data);
export const getAltSuppliers = (id) => api.get(`/products/alternative-suppliers?id=${id}`).then(res => res.data);
export const getCompanies = () => api.get('/companies').then(res => res.data);
export const getRiskAssessment = (id) => api.get(`/companies/risk-assessment?id=${id}`).then(res => res.data);
export const getOrders = () => api.get('/orders').then(res => res.data);
export const getSupplyPath = (id) => api.get(`/orders/supply-path?orderId=${id}`).then(res => res.data);
export const getCostBreakdown = (id) => api.get(`/orders/cost-breakdown?orderId=${id}`).then(res => res.data);
export const getHealth = () => api.get('/analytics/supply-chain-health').then(res => res.data);
export const getImpact = (id) => api.get(`/analytics/impact-analysis?supplier=${id}`).then(res => res.data);
export const getStockLevels = (id) => api.get(`/analytics/stock-levels?product=${id}`).then(res => res.data);
export const getForecastDelays = (months = 3) => api.get(`/analytics/forecast-delays?months=${months}`).then(res => res.data);
export const getOptimalRoute = (from, to) => api.get(`/routes/optimal?from=${from}&to=${to}`).then(res => res.data);
export const getLocations = () => api.get('/locations').then(res => res.data); //