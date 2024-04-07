import logo from './logo.svg';
import './App.css';
import {useEffect, useState} from "react";

function App() {
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    const eventSource = new EventSource('http://localhost:8080/populate');

    eventSource.onmessage = (event) => {
      setMessages((prevMessages) => [...prevMessages, JSON.parse(event.data)]);
    };

    setTimeout(() => eventSource.close(), 3000)

    return () => {
      eventSource.close();
    };
  }, []);

  return (
      <div className="App">
        <header className="App-header">
          <h1>Pokemons Capturados {messages.length}</h1>
          <ul style={{ padding: 0}}>
            {messages.map((message, index) => (
                <>
                {message?.sprites?.front_default &&
                <img title={message.name} src={message.sprites.front_default} width={64} height={64} />}
                </>
            ))
                }
          </ul>
        </header>
      </div>
  );
}

export default App;
