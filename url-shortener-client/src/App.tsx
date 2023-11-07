import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from './pages/home/Home';
import Docs from './pages/Docs/Docs';

const App = () => (
  <>
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='docs' element={<Docs />} />
      </Routes>
    </BrowserRouter>
  </>
);

export default App;