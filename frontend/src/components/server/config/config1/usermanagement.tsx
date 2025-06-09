import './usermanagement.css';

const UserManagement = () => {
  return (
    <div className="card">
      <h2>User Management</h2>
      <input type="text" placeholder="Add username" />
      <button>Add User</button>
      <ul>
        <li>admin</li>
        <li>devops</li>
        <li>analyst</li>
      </ul>
    </div>
  );
};

export default UserManagement;
