export default function StatusBadge({ status }) {
  const getColor = () => {
    switch (status) {
      case 'active': case 'delivered': return 'low';
      case 'pending': case 'in_transit': return 'medium';
      case 'discontinued': case 'delayed': return 'high';
      default: return 'medium';
    }
  };
  return <span className={`badge ${getColor()}`}>{status}</span>;
}