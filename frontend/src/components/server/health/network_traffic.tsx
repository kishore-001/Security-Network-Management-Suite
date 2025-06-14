import './network_traffic.css';
import wifiIcon from '../../assets/icons/wifi.png';

const NetworkTraffic: React.FC = () => {
  return (
    <div className="health-card">
      <div className="health-card-header">
        <div className="health-icon-container health-gradient-bg">
          <img src={wifiIcon} alt="Network Icon" className="health-cpu-img" />
        </div>
        <div>
          <h4>Network Traffic</h4>
          <p>Real-time data transfer</p>
        </div>
        <span className="health-expand">⤢</span>
      </div>
      <div className="health-network-traffic">
        <div>
          <span className="health-label" style={{ color: '#f97316' }}>Transmit</span>
          <h3>5.2 <span className="health-label">MB/s</span></h3>
        </div>
        <div>
          <span className="health-label" style={{ color: '#0ea5e9' }}>Receive</span>
          <h3>0.8 <span className="health-label">MB/s</span></h3>
        </div>
      </div>
    </div>
  );
};

export default NetworkTraffic;