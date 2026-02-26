export const formatCurrency = (value) => {
  if (value == null) return '$0.00';
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
    minimumFractionDigits: 0,
    maximumFractionDigits: 2,
  }).format(value);
};

export const formatPercent = (value) => {
  if (value == null) return '0%';
  return `${(value * 100).toFixed(0)}%`;
};

export const formatNumber = (value, decimals = 1) => {
  if (value == null) return '0';
  return Number(value).toFixed(decimals).replace(/\.0+$/, ''); 
};