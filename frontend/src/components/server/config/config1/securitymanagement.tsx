import { useState } from "react";
import "./securitymanagement.css";
import ModalWrapper from "./modalwrapper"; // adjust the path if needed

const SecurityManagement = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <>
      <div className="card security-card" onClick={() => setIsModalOpen(true)}>
        <div className="modal-title"></div>
      
        <h2> Security Management</h2>
        <p>Manage SSH keys and passwords</p>
      </div>

      {isModalOpen && (
        <ModalWrapper title="Security Management" onClose={() => setIsModalOpen(false)}>
          <p className="subtitle">
            Manage authentication credentials and security access controls
          </p>

          {/* SSH Key Section */}
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
                padding: "5px",
              }}
            ></textarea>
            <button type="button"  style={{ paddingLeft:"25px",marginTop: "5px",alignItems:"center",marginBottom:"5px", width:"500px", background: "#f1f5f9", color: "#1e293b" }}>
              🔑 Upload SSH Key
            </button>

            {/* Password Management */}
            <input type="text" placeholder="Enter username" />

            <input type="password" placeholder="Current password" />

            <input type="password" placeholder="New password" />

            <input type="password" placeholder="Confirm password" />

            <button type="submit" >📁 Update Credentials</button>
          </form>
        </ModalWrapper>
      )}
    </>
  );
};

export default SecurityManagement;
