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
    fetchRoute('letters', setLetters);
    fetchRoute('words', setWords);
  }, []);

  const handleChange = (event: any) => {
    event.preventDefault();
    setText(event.target.value);
    fetchRoute('letters', setLetters);
    fetchRoute('words', setWords);
  };

  const backendUrl =
    import.meta.env.VITE_APP_API_URL || 'http://localhost:8080';

  const fetchRoute = async (
    route: string,
    setState: React.Dispatch<React.SetStateAction<never[]>>
  ) => {
    try {
      const response = await fetch(`${backendUrl}/${route}`, {
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
      setState(result);
    } catch {
      throw new Error('Unexpected error');
    }
  };

  return (
    <>
      <div>
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            width: '100%',
            height: '300px',
            boxSizing: 'border-box',
          }}
        >
          <h3>Enter some arbitrary text, and then observe the frequencies.</h3>
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
            <h3>Letter frequencies</h3>
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
            <h3>Word frequencies</h3>
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
