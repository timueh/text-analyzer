import { useEffect, useState } from 'react';
import './App.css';
import { CartesianGrid, XAxis, YAxis, Tooltip, BarChart, Bar } from 'recharts';

function App() {
  const [data, setData] = useState([]);
  const [text, setText] = useState('');

  useEffect(() => {
    fetchData();
  }, []);

  const handleChange = (event: any) => {
    event.preventDefault();
    setText(event.target.value);
    fetchData();
  };

  const backendUrl =
    import.meta.env.VITE_APP_API_URL || 'http://localhost:8080';

  const fetchData = async () => {
    try {
      const response = await fetch(`${backendUrl}/zipf`, {
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
    } catch {
      throw new Error('Unexpected error');
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
