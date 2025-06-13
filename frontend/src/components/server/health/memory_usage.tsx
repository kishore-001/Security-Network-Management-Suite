import React from 'react';
import './memory_usage.css';
import memory from '../../../assets/icon/memory.svg'
const MemoryUsage: React.FC = () => {
  return (
    <div className="card memory-card">
      <div className="card-header">
       <div className="icon-container gradient">
          <img src={memory} alt="memory Icon" className="img" />
        </div>
        <div>
          <h4>Memory Usage</h4>
          <p>21.8GB / 32GB</p>
        </div>
        <span className="expand">⤢</span>
      </div>
      <div className="card-body">
        <h2 className="value green">68.2% <span className="label">used</span></h2>
        <div className="bar-track"><div className="bar-fill" style={{ width: '68.2%' }}></div></div>
        <div className="range">
          <span>Available: 10.2GB</span>
          <span>Used: 21.8GB</span>
        </div>
      </div>
    </div>
  );
};

export default MemoryUsage;
