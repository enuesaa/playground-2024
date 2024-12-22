import { Stagehand } from "@browserbasehq/stagehand";
import { z } from "zod";

const stagehand = new Stagehand({
  env: "LOCAL",
});

await stagehand.init();
await stagehand.page.goto("https://github.com/browserbase/stagehand");
await stagehand.act({ action: "click on the contributors" });
const contributor = await stagehand.extract({
  instruction: "extract the top contributor",
  schema: z.object({
    username: z.string(),
    url: z.string(),
  }),
});
await stagehand.close();
console.log(`Our favorite contributor is ${contributor.username}`);
