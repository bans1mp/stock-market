import "./index.css";
import { motion } from "framer-motion";
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";

const data = [
  { time: "10 AM", price: 150 },
  { time: "11 AM", price: 155 },
  { time: "12 PM", price: 160 },
  { time: "1 PM", price: 158 },
  { time: "2 PM", price: 162 },
];

function Card({ children, className }) {
  return <div className={`card ${className}`}>{children}</div>;
}

function CardContent({ children }) {
  return <div>{children}</div>;
}

function Button({ children, className, onClick }) {
  return (
    <button className={`button ${className}`} onClick={onClick}>
      {children}
    </button>
  );
}

function App() {
  return (
    <div className="container">
      <motion.h1
        className="title"
        initial={{ opacity: 0, y: -20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        Stock Tracker
      </motion.h1>
      <Card className="chart-card">
        <CardContent>
          <ResponsiveContainer width="100%" height={300}>
            <LineChart data={data}>
              <XAxis dataKey="time" stroke="#ddd" />
              <YAxis stroke="#ddd" />
              <Tooltip contentStyle={{ background: "#333", border: "none" }} />
              <Line type="monotone" dataKey="price" stroke="#00c6ff" strokeWidth={2} dot={{ r: 4 }} />
            </LineChart>
          </ResponsiveContainer>
          <div className="button-container">
            <Button className="buy-button">Buy</Button>
            <Button className="sell-button">Sell</Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

export default App;
