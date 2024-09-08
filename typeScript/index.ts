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
  .command("create <name> <note>")
  .description("create note ")
  .action(async (note: string, name: string) => {
    console.log(note, name);
    await noteServices.save(note, name);
  });

program
  .command("find")
  .description("find all note ")
  .action(async () => {
    const result = await noteServices.findAll();
    console.log(result);
  });

program.parse();
