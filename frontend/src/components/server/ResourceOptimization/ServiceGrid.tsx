import React from 'react';
import './serviceGrid.css';

const services = [
  'Web Servers', 'Application Servers', 'Database Services',
  'Monitoring Agents', 'Log Stack Components', 'Firewall & Network',
  'Backup Services', 'SSH/Remote Access', 'Cron', 'DNS/Cache'
];

const ServiceGrid: React.FC = () => {
  return (
    <div className="service-grid">
      {services.map(service => (
        <div key={service} className="service-box">
          <div className="service-icon">📦</div>
          <div className="service-label">{service}</div>
          <div className="service-status">🟢 running</div>
          <div className="service-meta">1 connection</div>
        </div>
      ))}
    </div>
  );
};

export default ServiceGrid;
