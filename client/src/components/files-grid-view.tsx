import { useEffect, useState } from "react";
import FileCard from "./file-card";
import type { FileDto } from "@/lib/types/file";

interface FilesGridViewProps { };

type FetchFilesResponse = {
  message: string,
  files: FileDto[],
}

const FilesGridView: React.FC<FilesGridViewProps> = ({ }) => {
  const [files, setFiles] = useState<FileDto[]>([]);

  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const res = await fetch("/api/files");
        if (!res.ok) {
          throw new Error("error fetching files");
        }

        const data: FetchFilesResponse = await res.json();
        console.log(data);

        setFiles(data.files ?? [])
      } catch (err) {
        console.error(err);
      }
    }
    fetchFiles();
  }, []);

  return (
    <div className="grid sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 px-8 gap-4 auto-rows-min">
      {files.map((f, i) => {
        return (
          <FileCard file={f} key={i} />
        );
      })}
    </div>
  );
}

export default FilesGridView;
