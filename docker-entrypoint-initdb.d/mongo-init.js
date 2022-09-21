db.createUser({
  user: "root",
  pwd: "root",
  roles: [
    {
      role: "readWrite",
      db: "websiteDev",
    },
  ],
});

db.createCollection("myinfoDev");
db.myinfoDev.insertMany([
  {
    lang: "en",
    image:
      "https://s.gravatar.com/avatar/1bf51742e09fee803505b9c0a845e262?s=80",
    name: "Edilberto Vazquez Luna",
    technologies: [
      "JavaScript",
      "TypeScript",
      "NextJS",
      "CSS",
      "SASS",
      "Graphql",
      "NodeJS",
      "Golang",
      "Docker",
      "Git",
      "Sequelize",
      "MongoDB",
      "SQL",
    ],
    projects: [
      {
        name: "Edilberto-Vazquez/website-ui",
        description: "Personal website",
        repository: "https://github.com/Edilberto-Vazquez/website-ui",
        technologies: ["JavaScript", "TypeScript", "NextJS", "React", "SASS"],
        website: "",
      },
    ],
    jobs: [
      {
        position: "Front-end",
        company: "Full Stack",
        location: "Puebla, Mexico",
        description:
          "Professional internship contract as Front-End developer in FullStack. Work done, migration of the ferreyarns website from Worpress to ReactJS, development of the Vetech site for the management of pet consultations for veterinarians.",
        dates: { from: "2021-05", to: "2021-09" },
      },
      {
        position: "Front-end",
        company: "Kannbal",
        location: "Mexico City, Mexico",
        description:
          "Front-End developer in kannbal conuslting. Development of web solutions for LMS like Edmodo, Moodle etc, using HTML, CSS, JavaScript, and Web Development with ReactJS.",
        dates: { from: "2022-01", to: "2022-02" },
      },
      {
        position: "Full-Stack",
        company: "BluePixel",
        location: "Mexico City, Mexico",
        description:
          "Front-end development with NextJS, Graphql, Apollo Client, Material-UI and Back-end development with NodeJS, ApolloServer, Python, MongoDB, MySQL and Docker",
        dates: { from: "2022-03", to: "present" },
      },
    ],
    resume:
      "I started in this world of programming in 2016 when I entered university, but during the pandemic I felt that university was not helping me achieve my goals, so I decided to study on my own at platzi. That was the place where I started to learn about JavaScript and Frontend technologies, but not only that, I also decided to explore other areas like Backend, some Data Science stuff and programming languages like Python, Golang and Rust. And thanks to all the courses I took at platzi and hours and hours of self-study, I was able to create frontend and backend applications. It also opened the doors for me to be part of the Platzi master community and face new challenges",
    description:
      "Hello, Welcome to my Website where you can see my projects and other things. I am currently working as a JavaScript developer but I work with other languages like Go, for example the backend of this site is made with GO",
  },
  {
    lang: "es",
    image:
      "https://s.gravatar.com/avatar/1bf51742e09fee803505b9c0a845e262?s=80",
    name: "Edilberto Vazquez Luna",
    technologies: [
      "JavaScript",
      "TypeScript",
      "NextJS",
      "CSS",
      "SASS",
      "Graphql",
      "NodeJS",
      "Golang",
      "Docker",
      "Git",
      "Sequelize",
      "MongoDB",
      "SQL",
    ],
    projects: [
      {
        name: "Edilberto-Vazquez/website-ui",
        description: "Sitio web personal",
        repository: "https://github.com/Edilberto-Vazquez/website-ui",
        technologies: ["JavaScript", "TypeScript", "NextJS", "React", "SASS"],
        website: "",
      },
    ],
    jobs: [
      {
        position: "Front-end",
        company: "Full Stack",
        location: "Puebla, México",
        description:
          "Contrato de prácticas profesionales como desarrollador Front-End en FullStack. Trabajo realizado, migración del sitio web de ferreyarns de Worpress a ReactJS, desarrollo del sitio Vetech para la gestión de consultas de mascotas para veterinarios.",
        dates: { from: "2021-05", to: "2021-09" },
      },
      {
        position: "Front-end",
        company: "Kannbal",
        location: "Ciudad de México, México",
        description:
          "Desarrollador Front-End en kannbal. Desarrollo de soluciones web para LMS como Edmodo, Moodle, etc, utilizando HTML, CSS, JavaScript y desarrollo web con ReactJS.",
        dates: { from: "2022-01", to: "2022-02" },
      },
      {
        position: "Full-Stack",
        company: "BluePixel",
        location: "Ciudad de México, México",
        description:
          "Desarrollo front-end con NextJS, Graphql, Apollo Client, Material-UI y desarrollo Back-end con NodeJS, ApolloServer, Python, MongoDB, MySQL y Docker.",
        dates: { from: "2022-03", to: "actual" },
      },
    ],
    resume:
      "Empecé en este mundo de la programación en 2016 cuando entré a la universidad, pero durante la pandemia sentí que la universidad no me ayudaba a lograr mis metas, así que decidí estudiar por mi cuenta en platzi. Ese fue el lugar donde comencé a aprender sobre tecnologías JavaScript y Frontend, pero no solo eso, también decidí explorar otras áreas como Backend, algunas cosas de Data Science y lenguajes de programación como Python, Golang y Rust. Y gracias a todos los cursos que tomé en platzi y horas y horas de autoaprendizaje, pude crear aplicaciones frontend y backend. Ademas me abrio las puertas para formar parte de la comunidad de Platzi master y afrontar nuevos retos",
    description:
      "Hola, Bienvenidos a mi Sitio Web donde pueden ver mis proyectos y otras cosas. Actualmente estoy trabajando como desarrollador de JavaScript pero trabajo con otros lenguajes como Go, por ejemplo el backend de este sitio esta echo con GO",
  },
]);
