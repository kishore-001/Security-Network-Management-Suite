import React from 'react';
import './cpu_usage.css';
import cpu from '../../../assets/icon/cpu.svg'; // ✅ Use your actual icon path

const CpuUsage: React.FC = () => {
  return (
    <div className="card cpu-card">
      <div className="card-header">
        <div className="icon-container gradient-bg">
          <img src={cpu} alt="CPU Icon" className="img" />
        </div>
        <div>
          <h4>CPU Usage</h4>
          <p>Current load</p>
        </div>
        <span className="expand">⤢</span>
      </div>
      <div className="card-body">
        <h2 className="value green">8.5% <span className="label">current</span></h2>
        <div className="bar-track"><div className="bar-fill" style={{ width: '8.5%' }}></div></div>
        <div className="range">
          <span>Avg: 12.3%</span>
          <span>Peak: 45.2%</span>
        </div>
      </div>
    </div>
  );
};

export default CpuUsage;
