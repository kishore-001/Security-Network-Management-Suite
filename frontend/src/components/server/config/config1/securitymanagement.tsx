import { useState } from "react";
import "./securitymanagement.css";
import ModalWrapper from "./modalwrapper"; // adjust path if needed

const SecurityManagement = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <>
      <div className="card security-card" onClick={() => setIsModalOpen(true)}>
        <div className="security-icon">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            height="24"
            viewBox="0 0 24 24"
            width="24"
            fill="#ffffff"
          >
            <path d="M0 0h24v24H0z" fill="none" />
            <path d="M12 1L3 5v6c0 5.25 3.75 10.74 9 12 5.25-1.26 9-6.75 9-12V5l-9-4zM12 17c-2.76 0-5-2.24-5-5 0-1.38.56-2.63 1.46-3.54C9.37 7.56 10.62 7 12 7s2.63.56 3.54 1.46C16.44 9.37 17 10.62 17 12c0 2.76-2.24 5-5 5z" />
          </svg>
        </div>
        <h3>Security Management</h3>
        <p>Manage SSH keys and passwords</p>
      </div>

      {isModalOpen && (
        <ModalWrapper
          title="Security Management"
          onClose={() => setIsModalOpen(false)}
        >
          <p className="subtitle">
            Manage authentication credentials and security access controls
          </p>

          <form>
            <label>SSH Public Key Management</label>
            <textarea
              placeholder="Paste SSH public key here..."
              style={{
                width: "80%",
                height: "80px",
                backgroundColor: "#0f172a",
                color: "white",
                border: "1px solid #475569",
                borderRadius: "8px",
                padding: "10px",
              }}
            ></textarea>

            <button
              type="button"
              style={{
                marginTop: "5px",
                width: "500px",
                background: "#f1f5f9",
                color: "#1e293b",
              }}
            >
              🔑 Upload SSH Key
            </button>

            <input type="text" placeholder="Enter username" />
            <input type="password" placeholder="Current password" />
            <input type="password" placeholder="New password" />
            <input type="password" placeholder="Confirm password" />

            <button type="submit">📁 Update Credentials</button>
          </form>
        </ModalWrapper>
      )}
    </>
  );
};

export default SecurityManagement;
