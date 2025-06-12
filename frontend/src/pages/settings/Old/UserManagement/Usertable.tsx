import React from 'react';
import './UserManagement.css';

const UserTable: React.FC = () => {
  return (
    <div className="table-wrapper">
      <h3>User Management</h3>
      <p>Allowed Users</p>
      <table>
        <thead>
          <tr>
            <th>S.No</th>
            <th>Username</th>
            <th>Role</th>
            <th>Email Address</th>
          </tr>
        </thead>
        <tbody>
          {/* rows go here */}
        </tbody>
      </table>
    </div>
  );
};

export default UserTable;
