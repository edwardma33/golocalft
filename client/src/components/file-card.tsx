import { Card, CardAction, CardDescription, CardFooter, CardHeader, CardTitle } from "./ui/card";
import { Button } from "./ui/button";
import { Download, Eye, Trash } from "lucide-react";
import type { FileDto } from "@/lib/types/file";
import { useState } from "react";

interface FileCardProps {
  file: FileDto
};

type DeleteResponse = {
  message: string,
};

const FileCard: React.FC<FileCardProps> = ({ file }) => {
  const [isDeleted, setIsDeleted] = useState<boolean>(false);
  const handleDelete = async () => {
    try {
      const res = await fetch(`/api/delete/${file.id}`);

      if (!res.ok) throw new Error("error delete id: " + file.id);

      const data: DeleteResponse = await res.json();

      console.log(data.message);
      setIsDeleted(true);
      // window.location.reload();
    } catch (err) {
      console.log(err);
    }
  }
  
  const handleDownload = async () => {
    try {
      const res = await fetch(`/api/download/${file.id}`);
      if (!res.ok) throw new Error("error download id: " + file.id);

      const blob = await res.blob();

      const objectUrl = window.URL.createObjectURL(blob);

      const link = document.createElement("a");
      link.href = objectUrl;
      link.download = file.name + file.extension;
      document.body.appendChild(link);
      link.click();

      link.remove();
      window.URL.revokeObjectURL(objectUrl);
    } catch (err) {
      console.log(err);
    }
  }

  return (
    <Card className={isDeleted ? "hidden" : ""} >
      <CardHeader>
        <CardTitle>{file.name}{file.extension}</CardTitle>
        <CardDescription>{file.createdAt.toString()}</CardDescription>
        <CardAction>
          <Button size={"icon"}>
            <Eye />
          </Button>
        </CardAction>
      </CardHeader>
      <CardFooter className="grid grid-cols-2 gap-2">
        <Button onClick={handleDownload}>
          <Download />
          <span>Save</span>
        </Button>
        <Button onClick={handleDelete} variant={"destructive"}>
          <Trash />
          <span>Delete</span>
        </Button>
      </CardFooter>
    </Card>
  );
}

export default FileCard;
