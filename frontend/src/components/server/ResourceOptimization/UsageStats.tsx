import React from 'react';
import './usageStats.css';

interface UsageStatsProps {
  label: string;
  value: number;
}

const UsageStats: React.FC<UsageStatsProps> = ({ label, value }) => {
  return (
    <div className="usage-stat">
      <div className="circle">
        <span>{value}%</span>
      </div>
      <p>{label}</p>
    </div>
  );
};

export default UsageStats;
