import { initializeApp, cert } from "firebase-admin/app";
import { getStorage } from "firebase-admin/storage";
import dotenv from "dotenv";
import tar from "tar";

tar
  .c(
    {
      file: "specs.tar",
    },
    ["./specs"]
  )
  .then(() => {
    dotenv.config();

    initializeApp({
      credential: cert({
        projectId: process.env.FIREBASE_PROJECT_ID,
        privateKey: (process.env.FIREBASE_PRIVATE_KEY as string).replace(
          /\\n/g,
          "\n"
        ),
        clientEmail: process.env.FIREBASE_CLIENT_EMAIL,
      }),
      storageBucket: process.env.FIREBASE_STORAGE_BUCKET,
    });

    const bucket = getStorage().bucket();

    bucket
      .upload("./specs.tar", {
        public: true,
      })
      .then(() => {
        console.log("Uploaded");
      })
      .catch((err) => {
        console.error(err);
      });
  })
  .catch((err) => {
    console.error(err);
  });
