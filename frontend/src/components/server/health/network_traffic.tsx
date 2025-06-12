import React from 'react';
import './network_traffic.css';
import network from './../../../assets/icon/network.svg';
const NetworkTraffic: React.FC = () => {
  return (
    <div className="card network-card">
      <div className="card-header">
<div className="icon-container gradient_network">
          <img src={network} alt="network Icon" className="img" />
        </div>        <div>
          <h4>Network Traffic</h4>
          <p>Real-time data transfer</p>
        </div>
        <span className="expand">⤢</span>
      </div>
      <div className="card-body traffic-body">
        <div className="traffic-row">
          <span className="dot orange" /> Transmit <span className="rate">5.2 MB/s</span>
        </div>
        <div className="traffic-row">
          <span className="dot blue" /> Receive <span className="rate">0.8 MB/s</span>
        </div>
      </div>
    </div>
  );
};

export default NetworkTraffic;
