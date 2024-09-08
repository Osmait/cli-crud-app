import ollama from "ollama";
import { Note } from "./note.model";
import { NoteRepository } from "./repository";

export class NoteServices {
  private noteRepository: NoteRepository;

  constructor(noteRepository: NoteRepository) {
    this.noteRepository = noteRepository;
  }
  async save(note: string, name: string) {
    const noteToSave: Note = {
      id: Math.random().toString(),
      name: name,
      description: note,
      createdAt: Date.now().toString(),
      updateAt: Date.now().toString(),
    };
    await this.noteRepository.save(noteToSave);
  }

  async findAll(): Promise<Note[]> {
    return await this.noteRepository.findAll();
  }

  async summary() {
    const notes = await this.noteRepository.findAll();
    const description = notes.map((n) => n.description);

    const response = await ollama.chat({
      model: "llama3.1",
      messages: [
        {
          role: "user",
          content: `From the descriptions that I am going to give you, tell me which ones you think are the most important.${description} Just give me the grades, don't tell me why you chose them. numeric order `,
        },
      ],
    });
    console.log(response.message.content);
  }
}
