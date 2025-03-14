import "./index.css";
import homeImage from "./assets/home.jpg";
import Navbar from "./components/Navbar.jsx";
import Footer from "./components/Footer.jsx";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

function App() {
  const navigate = useNavigate();
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem("token");
    setIsLoggedIn(!!token);
  }, []);

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
      <div className="container flex-grow-1 d-flex align-items-center justify-content-between mt-5 px-5">
        <div className="text-section" style={{ maxWidth: "50%", opacity: "0.75", color: "rgba(200,200,200,0.85)", textShadow: "2px 2px 6px rgba(0,0,0,0.8)" }}>
          <h1 className="fw-bold display-2 text-success border-start border-4 border-success ps-3" style={{ fontFamily: "Poppins, sans-serif" }}>
            Welcome to <br />
            <span className="text-danger">StockSim</span>
          </h1>
          <p className="lead fs-2">
            <span className="fw-bold text-success">Simulate</span>, <span className="fw-bold text-success">Trade</span>, <br />
            and <span className="fw-bold text-success">Master the Market</span> <br />
            with Realistic Stock Exchange Models.
          </p>
        </div>
        <div className="button-section d-flex align-items-center gap-4">
          {isLoggedIn ? (
            <button 
              className="btn px-5 py-4 fs-3 fw-bold border border-white text-white" 
              style={{ 
                minWidth: "220px", 
                backgroundColor: "transparent", 
                opacity: "0.75",
                transition: "0.3s ease-in-out",
              }}
              onMouseOver={(e) => {
                e.target.style.backgroundColor = "rgba(255, 255, 255, 0.2)";
                e.target.style.color = "#000"; // Change text color to black
                e.target.style.opacity = "1";
              }}
              onMouseOut={(e) => {
                e.target.style.backgroundColor = "transparent";
                e.target.style.color = "white";
                e.target.style.opacity = "0.75";
              }}
              onClick={() => navigate("/stocks")}
            >
              Start Trading
            </button>
          ) : (
            <>
              <button 
                className="btn btn-success px-5 py-4 fs-3 fw-bold" 
                style={{ 
                  minWidth: "220px", 
                  opacity: "0.85",
                  transition: "0.3s ease-in-out",
                }}
                onMouseOver={(e) => {
                  e.target.style.backgroundColor = "#157347"; // Darker green
                  e.target.style.color = "#ddd"; // Lighter text
                }}
                onMouseOut={(e) => {
                  e.target.style.backgroundColor = "#198754"; // Original green
                  e.target.style.color = "white";
                }}
                onClick={() => navigate("/login")}
              >
                Login
              </button>
              <button 
                className="btn btn-outline-danger px-5 py-4 fs-3 fw-bold" 
                style={{ 
                  minWidth: "220px", 
                  opacity: "0.85",
                  transition: "0.3s ease-in-out",
                }}
                onMouseOver={(e) => {
                  e.target.style.backgroundColor = "rgba(220, 53, 69, 0.85)"; // Red with opacity
                  e.target.style.color = "#fff"; // Ensure white text
                }}
                onMouseOut={(e) => {
                  e.target.style.backgroundColor = "transparent"; // Back to outline style
                  e.target.style.color = "#dc3545"; // Red text
                }}
                onClick={() => navigate("/register")}
              >
                Register
              </button>
            </>
          )}
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default App;