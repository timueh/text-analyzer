import { useEffect, useState } from 'react';
import './App.css';
import {
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  BarChart,
  Bar,
  ResponsiveContainer,
} from 'recharts';

function App() {
  const [letters, setLetters] = useState([]);
  const [words, setWords] = useState([]);
  const [text, setText] = useState('');

  useEffect(() => {
    fetchLetters();
  }, []);

  const handleChange = (event: any) => {
    event.preventDefault();
    setText(event.target.value);
    fetchLetters();
    fetchWords();
  };

  const backendUrl =
    import.meta.env.VITE_APP_API_URL || 'http://localhost:8080';

  const fetchLetters = async () => {
    try {
      const response = await fetch(`${backendUrl}/letters`, {
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
      setLetters(result);
    } catch {
      throw new Error('Unexpected error');
    }
  };

  const fetchWords = async () => {
    try {
      const response = await fetch(`${backendUrl}/words`, {
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
      setWords(result);
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
            flexDirection: 'column',
            width: '100%',
            height: '300px', // Adjust height as needed
            boxSizing: 'border-box',
            // border: '1px solid black',
          }}
        >
          <div style={{ flex: 1, boxSizing: 'border-box', padding: '10px' }}>
            <textarea
              style={{
                width: '100%', // Make textarea span the full width of its parent
                height: '100px', // Example height, adjust as needed
              }}
              value={text}
              placeholder="Analyze this text"
              onChange={handleChange}
            />
          </div>
          <div style={{ flex: 1, boxSizing: 'border-box', padding: '10px' }}>
            <ResponsiveContainer width="100%" height={200}>
              <BarChart width={600} height={300} data={letters}>
                <CartesianGrid />
                <XAxis dataKey="name" />
                <YAxis allowDecimals={false} />
                <Tooltip />
                <Bar dataKey="value" fill="#8884d8" />
              </BarChart>
            </ResponsiveContainer>
          </div>
          <div style={{ flex: 1 }}>
            <ResponsiveContainer width="100%" height={200}>
              <BarChart width={600} height={300} data={words}>
                <CartesianGrid />
                <XAxis dataKey="name" />
                <YAxis allowDecimals={false} />
                <Tooltip />
                <Bar dataKey="value" fill="#8884d8" />
              </BarChart>
            </ResponsiveContainer>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
