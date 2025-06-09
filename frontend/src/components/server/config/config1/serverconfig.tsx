import './serverconfig.css';

const ServerConfig = () => {
  return (
    <div className="card">
      <h2>Server Configuration</h2>
      <input type="text" placeholder="Enter config path..." />
      <textarea placeholder="Paste configuration data..." />
      <button>Apply Configuration</button>
    </div>
  );
};

export default ServerConfig;
