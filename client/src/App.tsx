import FilesGridView from "./components/files-grid-view";
import Navbar from "./components/navbar";
import "./index.css";

export function App() {
  return (
    <div className="w-screen h-screen flex-col">
      <Navbar />
      <FilesGridView />
    </div>
  );
}

export default App;
