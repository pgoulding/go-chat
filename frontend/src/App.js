import './App.css';
import Header from "./components/Header/"
import ChatHistory from './components/ChatHistory';

function App() {
  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={[]}/>
      <button onClick={null} >Send</button>
    </div>
  );
}

export default App;
