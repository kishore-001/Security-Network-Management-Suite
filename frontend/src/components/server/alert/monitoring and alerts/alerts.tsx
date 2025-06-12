import React from 'react';
import './alerts.css';

const Alerts: React.FC = () => {
  return (
    <div className="alerts-card">
      <h3>Alerts</h3>
      <div className="circle-alert">
        <div className="circle">1551</div>
        <span>Total</span>
      </div>
      <div className="alert-counts">
        <div className="red">1401 High</div>
        <div className="yellow">39 Medium</div>
        <div className="blue">111 Low</div>
      </div>
    </div>
  );
};

export default Alerts;
