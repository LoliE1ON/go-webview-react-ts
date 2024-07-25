import './App.css'

function App() {
    const sendMessageToGo = () => {
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-expect-error
        window.chrome.webview.postMessage(JSON.stringify({ message: "sendMessage"}));
    };

  return (
      <>
          <button onClick={sendMessageToGo}>
              Exit
          </button>
      </>
  )
}

export default App
