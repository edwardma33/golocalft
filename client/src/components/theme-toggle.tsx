import { Moon, Sun } from "lucide-react";
import { Button } from "./ui/button";
import { useEffect, useState } from "react";

interface ThemeToggleProps { };

const ThemeToggle: React.FC<ThemeToggleProps> = ({ }) => {
  const [isDark, setIsDark] = useState<boolean>(false);

  const toggleTheme = () => {
    const newIsDark = !isDark;
    setIsDark(newIsDark);
    document.documentElement.classList.toggle("dark", newIsDark);
    localStorage.setItem("theme", newIsDark ? "dark" : "light");
  };

  useEffect(() => {
    const saved = localStorage.getItem("theme");
    const isDark = saved === "dark";
    setIsDark(isDark);
    document.documentElement.classList.toggle("dark", isDark);
  }, []);

  return (
    <Button size={"icon"} variant={"ghost"} onClick={toggleTheme}>
      <Sun className={isDark ? "absolute" : "hidden"} />
      <Moon className={isDark ? "hidden" : "absolute"} />
    </Button>
  );
}

export default ThemeToggle;
