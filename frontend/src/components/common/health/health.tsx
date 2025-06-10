// src/components/common/health/health.tsx
import React from 'react';
import './health.css';
import { FiTrendingUp, FiGlobe, FiShield, FiActivity } from 'react-icons/fi';

const Health: React.FC = () => {
  const stats = [
    {
      icon: <FiTrendingUp />,
      label: 'Total Bandwidth',
      value: '1.2 Gbps',
      change: '+12%',
      color: 'green',
    },
    {
      icon: <FiGlobe />,
      label: 'Active Connections',
      value: '446',
      change: '+5%',
      color: 'blue',
    },
    {
      icon: <FiShield />,
      label: 'Security Events',
      value: '23',
      change: '-8%',
      color: 'orange',
    },
    {
      icon: <FiActivity />,
      label: 'System Health',
      value: '98.5%',
      change: '+2%',
      color: 'green',
    },
  ];

  return (
    <div className="health-container">
      {stats.map((stat, index) => (
        <div className={`stat-card ${stat.color}`} key={index}>
          <div className="icon">{stat.icon}</div>
          <div className="value">{stat.value}</div>
          <div className="label">{stat.label}</div>
          <div className={`change ${stat.change.startsWith('-') ? 'negative' : 'positive'}`}>
            {stat.change}
          </div>
        </div>
      ))}
    </div>
  );
};

export default Health;