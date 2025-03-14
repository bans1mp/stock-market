import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import homeImage from "../assets/home.jpg";

function Register() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (password !== confirmPassword) {
      alert("Passwords do not match!");
      return;
    }

    try {
      const response = await axios.post("http://localhost:8080/register", {
        email,
        password,
      });

      if (response.status === 200) {
        console.log("Registration successful");
        navigate("/"); // Redirect to home after registration
      }
    } catch (error) {
      console.error("Error during registration:", error);
    }
  };

  return (
    <div 
      className="d-flex flex-column min-vh-100 text-white" 
      style={{ 
        backgroundImage: `url(${homeImage})`, 
        backgroundSize: "cover", 
        backgroundPosition: "center", 
        backgroundRepeat: "no-repeat", 
        backgroundBlendMode: "overlay",
        backgroundColor: "rgba(0, 0, 0, 0.85)" 
      }}
    >
      <Navbar />
      <div className="container d-flex flex-grow-1 justify-content-center align-items-center">
        <div className="card p-5 text-white" style={{ background: "rgba(30, 30, 30, 0.9)", borderRadius: "15px", width: "400px" }}>
          <h2 className="text-center mb-4 text-success">Register</h2>
          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label className="form-label">Email</label>
              <input 
                type="email" 
                className="form-control p-3" 
                value={email} 
                onChange={(e) => setEmail(e.target.value)} 
                required
              />
            </div>
            <div className="mb-4">
              <label className="form-label">Password</label>
              <input 
                type="password" 
                className="form-control p-3" 
                value={password} 
                onChange={(e) => setPassword(e.target.value)} 
                required
              />
            </div>
            <div className="mb-4">
              <label className="form-label">Confirm Password</label>
              <input 
                type="password" 
                className="form-control p-3" 
                value={confirmPassword} 
                onChange={(e) => setConfirmPassword(e.target.value)} 
                required
              />
            </div>
            <button type="submit" className="btn btn-danger w-100 fw-bold py-3">Register</button>
          </form>
          <p className="text-center mt-4">
            Already have an account? <a href="/login" className="text-danger fw-bold">Login</a>
          </p>
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default Register;
