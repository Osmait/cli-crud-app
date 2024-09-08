import { Command } from "commander";
import { version } from "./package.json";
import { NoteServices } from "./services";
import { NoteRepository } from "./repository";

const program = new Command();
const noteRepository = new NoteRepository();
const noteServices = new NoteServices(noteRepository);

program
  .version(version)
  .description("Una aplicación CLI de ejemplo con Commander");

program
  .command("create <note>")
  .description("create note ")
  .action(async (note: string) => {
    await noteServices.save(note);
  });

program
  .command("find")
  .description("find all note ")
  .action(async () => {
    const result = await noteServices.findAll();
    console.log(result);
  });

program.parse();
