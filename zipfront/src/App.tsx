import { useEffect, useState } from 'react';
import './App.css';
import { CartesianGrid, XAxis, YAxis, Tooltip, BarChart, Bar } from 'recharts';

function App() {
  const [data, setData] = useState([]);
  const [text, setText] = useState('');

  useEffect(() => {
    fetchData();
  }, []);

  const handleChange = (event) => {
    event.preventDefault();
    setText(event.target.value);
    fetchData();
  };

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:8080/zipf', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
        body: JSON.stringify({ data: text }),
      });

      if (!response.ok) {
        throw new Error('response was not ok');
      }

      const result = await response.json();
      setData(result);
    } catch (error) {
      throw new Error(error);
    }
  };

  return (
    <>
      <div>
        <h2>
          Enter arbitrary text, and observe the distribution of the letter
          frequencies.
        </h2>
        <div
          style={{
            display: 'flex',
          }}
        >
          <textarea
            value={text}
            placeholder="Analyze this text"
            onChange={handleChange}
          />
          <BarChart width={600} height={300} data={data}>
            <CartesianGrid />
            <XAxis dataKey="name" />
            <YAxis allowDecimals={false} />
            <Tooltip />
            <Bar dataKey="value" fill="#8884d8" />
          </BarChart>
        </div>
      </div>
    </>
  );
}

export default App;
