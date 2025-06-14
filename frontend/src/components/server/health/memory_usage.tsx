import React from 'react';
import './memory_usage.css';
//import memoryIcon from '../../assets/icons/memory.png';

const MemoryUsage: React.FC = () => {
  return (
    <div className="health-card">
      <div className="health-card-header">
        <div className="health-icon-container health-gradient-bg">
       </div>
        <div>
          <h4>Memory Usage</h4>
          <p>21.8GB / 32GB</p>
        </div>
        <span className="health-expand">⤢</span>
      </div>
      <h2 className="health-value health-green">68.2% <span className="health-label">used</span></h2>
      <div className="health-bar-track">
        <div className="health-bar-fill" style={{ width: '68.2%' }}></div>
      </div>
      <div className="health-range">
        <span>Available: 10.2GB</span>
        <span>Used: 21.8GB</span>
      </div>
    </div>
  );
};

export default MemoryUsage;