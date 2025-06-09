import './securitymanagement.css';

const SecurityManagement = () => {
  return (
    <div className="card">
      <h2>Security Management</h2>
      <div>
        <label>SSH Public Key Management</label>
        <textarea placeholder="Paste SSH public key here..." />
        <button className="upload-btn" disabled>Upload SSH Key</button>
      </div>
      <div>
        <label>Password Management</label>
        <input type="text" placeholder="Username" />
        <input type="password" placeholder="Current password" />
        <input type="password" placeholder="New password" />
        <input type="password" placeholder="Confirm password" />
        <button className="update-btn">Update Credentials</button>
      </div>
    </div>
  );
};

export default SecurityManagement;
