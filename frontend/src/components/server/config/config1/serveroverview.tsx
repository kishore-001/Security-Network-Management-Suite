import './serveroverview.css';

const ServerOverview = () => {
  return (
    <div className="card">
      <h2>Server Overview</h2>
      <ul>
        <li>Status: Online</li>
        <li>CPU Usage: 35%</li>
        <li>Memory: 6.2GB / 16GB</li>
      </ul>
    </div>
  );
};

export default ServerOverview;
