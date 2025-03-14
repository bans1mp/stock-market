import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import homeImage from "../assets/home.jpg";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("http://localhost:8080/login", { email, password });
      localStorage.setItem("token", response.data.token);
      navigate("/"); // Redirect to home after login
    } catch (error) {
      setError(error.response?.data?.message || "Login failed");
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
          <h2 className="text-center mb-4 text-success">Login</h2>
          {error && <p className="text-danger text-center">{error}</p>}
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
            <button type="submit" className="btn btn-danger w-100 fw-bold py-3">Login</button>
          </form>
          <p className="text-center mt-4">
            Don't have an account? <a href="/register" className="text-danger fw-bold">Register</a>
          </p>
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default Login;
