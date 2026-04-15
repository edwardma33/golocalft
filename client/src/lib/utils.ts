import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatDate(d: Date): string {
  const months = [
    "Jan", "Feb", "Mar", "Apr",
    "May", "Jun", "July", "Aug",
    "Sep", "Oct", "Nov", "Dev",
  ];
  const month = months[d.getMonth()];

  let hours = d.getHours() % 12;

  const meridiem = d.getHours() >= 12 ? "PM" : "AM";

  return `${month} ${d.getDate()} ${d.getFullYear()} ${hours == 0 ? 12 : hours}:${d.getMinutes()}${meridiem}`;
}
