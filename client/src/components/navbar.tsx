import { Upload } from "lucide-react";
import { Button } from "./ui/button";
import ThemeToggle from "./theme-toggle";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "./ui/dialog";
import { Input } from "./ui/input";
import type { SubmitEvent } from "react";
import { Field } from "./ui/field";
import { Label } from "./ui/label";

interface NavbarProps {};

const Navbar: React.FC<NavbarProps> = ({}) => {
  const handleUpload = async (e: SubmitEvent<HTMLFormElement>) => {
    console.log("uploading");
    e.preventDefault();

    try {
      const uploadForm = e.target;
      const formData = new FormData(uploadForm);

      const res = await fetch("/api/upload", {
        method: "POST",
        body: formData,
      });
      if (!res.ok) {
        throw new Error("upload error status: " + res.status);
      }

      const data = await res.json();

      window.location.reload();
      console.log(data);
    } catch (err: any) {
      console.error(err.message || err);
    }
  }

  return (
    <nav className="p-8 flex justify-between">
      <span className="text-2xl font-bold">LocalFT</span>
      <div className="flex gap-2 items-center">
        <Dialog>
          <DialogTrigger asChild>
            <Button type="button">
              <Upload />
              <span>Upload</span>
            </Button>
          </DialogTrigger>
          <form onSubmit={(e) => handleUpload(e)} id="upload-form">
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Upload a File</DialogTitle>
                <DialogDescription>Pick a file. Save it to LocalFT.</DialogDescription>
              </DialogHeader>
                <Field>
                  <Label htmlFor="file-name">Name</Label>
                  <Input form="upload-form" name="file-name" type="text" placeholder="optional custom name" />
                </Field>
                <Field>
                  <Label htmlFor="file">File</Label>
                  <Input form="upload-form" name="file" type="file" />
                </Field>
              <DialogFooter>
                <DialogClose asChild>
                  <Button type="button" variant="outline">Cancel</Button>
                </DialogClose>
                <Button form="upload-form" type="submit" >Upload File</Button>
              </DialogFooter>
            </DialogContent>
          </form>
        </Dialog>
        <ThemeToggle />
      </div>
    </nav>
  );
}

export default Navbar;
