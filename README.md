# Gemini
This repository provides a gRPC streaming service that queries the [Gemini](https://deepmind.google/technologies/gemini/#introduction) model.

## Pre-requisites
1. Protoc - https://grpc.io/docs/protoc-installation/
2. The application authenticates with Google cloud platform via the `GOOGLE_APPLICATION_CREDENTIALS` environment variable. When deployed, the cloud run instance automatically has access to the credentials. When developing locally you will need to set the variable. See https://cloud.google.com/docs/authentication/application-default-credentials
3. When changes are pushed to the main branch, [Cloud build](https://cloud.google.com/build?hl=en) will automatically build the code and deploy it into [Cloud run](https://cloud.google.com/run?hl=en).

## Documentation
1. Protocol Buffers - https://protobuf.dev/programming-guides/proto3
2. gRPC - https://grpc.io/docs/languages/go/quickstart/

## Guide
Inside of the root project directory you can run the below command to generate the server/client code for the relevant proto file.
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./protocol_buffers/gemini_service/gemini_service.proto    
```
Under the `./gemini_service` folder you will find the Gemini gRPC service that implements the server stubs that were generated with the above command. 

## Usage

You can query this gRPC service by sending a query to it such as `Tell me everything you know about the Roman empire` and Gemini will stream back a response that looks something like the following (once the stream has completed)

1. **Origins:**
   - The Roman Empire originated from the city of Rome, which was founded in 753 BC.
   - The early Roman kingdom was ruled by seven kings, and the last king, Tarquin the Proud, was overthrown in 509 BC, leading to the establishment of the Roman Republic.

2. **Republic:**
   - The Roman Republic was a period of political and military expansion, lasting from 509 BC to 27 BC.
   - The Republic was governed by the Senate, a body of wealthy and influential citizens, and annual elections were held to elect consuls, who served as the chief executives of the state.

3. **Conquests:**
   - The Romans embarked on a series of conquests, expanding their territory throughout Italy and beyond. They conquered Greece, North Africa, parts of the Middle East, and large areas of Europe.

4. **Military:**
   - The Roman army was one of the most powerful and disciplined in the ancient world, known for its strict organization, tactics, and engineering skills.
   - The army was divided into legions, consisting of approximately 5,000 soldiers, and was supported by auxiliary troops.

5. **Infrastructure:**
   - The Romans were skilled builders and engineers. They constructed extensive road networks, bridges, aqueducts, and public buildings throughout their empire.
   - These infrastructure projects facilitated communication, transportation, and trade.

6. **Government and Law:**
   - The Roman Republic was governed by a complex system of laws and institutions. The Senate held legislative power, and the assemblies of Roman citizens had the power to pass laws and elect officials.
   - The Roman legal system was highly developed and influenced many later legal systems.

7. **Culture:**
   - Roman culture was influenced by Greek, Etruscan, and other civilizations. They adopted many aspects of Greek art, literature, philosophy, and religion.
   - The Romans made significant contributions to art, architecture, and literature, creating a lasting legacy that continues to inspire modern societies.

8. **Religion:**
   - The Romans were polytheistic, worshipping a pantheon of gods and goddesses. Their religious practices included rituals, sacrifices, and festivals.
   - Later, the Roman Empire adopted Christianity as its official religion, leading to the spread of Christianity throughout Europe and beyond.

9. **Trade and Economy:**
   - The Roman Empire was a hub of trade and commerce, spanning vast areas of land and connecting diverse cultures.
   - Trade routes stretched across the Mediterranean Sea, connecting Europe, Asia, and Africa.
   - The economy was based on agriculture, with wealthy landowners owning large estates worked by slaves or tenant farmers.

10. **Decline and Fall:**
    - In the late 3rd century CE, the Roman Empire began to experience economic, political, and military crises.
    - The empire was divided into western and eastern halves, and the western half eventually fell under the control of barbarian tribes in the 5th century CE.
    - While the eastern half, known as the Byzantine Empire, survived until the 15th century CE, the fall of the western Roman Empire marked the end of the ancient world.
