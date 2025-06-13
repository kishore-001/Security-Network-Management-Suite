import React from 'react';
import './UserManagement.css';

interface User {
  id: number;
  name: string;
  email: string;
  role: string;
}

interface Props {
  users: User[];
}

const UserTable: React.FC<Props> = ({ users }) => {
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
          {users.map((user, index) => (
            <tr key={user.id}>
              <td>{index + 1}</td>
              <td>{user.name}</td>
              <td>{user.role}</td>
              <td>{user.email}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default UserTable;
