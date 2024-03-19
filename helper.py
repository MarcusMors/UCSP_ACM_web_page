def main():
    contactos = ["958799615",
    "945512169",
    "987339517",
    "970019510",
    "955174379",
    "964240030",
    "934407525",
    "954465862",
    "970898867",
    "29678316",
    "970019510"
    ]


    apoderados = ["María Lourdes Pérez Bedoya",
    "Anngie paredes",
    "María Paola Zerpa Huajardo",
    "Jimena Sofia Olmedo",
    "Jannya Paola Cáceres Rocha",
    "Carlos Eduardo sotelo pinto",
    "Eva Diaz Rafael",
    "Miguel Angel Urday Quintanilla",
    "Gleny  Núñez Subia de Cárdenas",
    "Jimena Sofia Olmedo"
    ]


    student_names = ["Janlennart Chávez Pérez",
    "Matt Duana Paredes",
    "Santiago Gómez Zerpa",
    "Thayra Kcana Olmedo",
    "Raúl Salazar Cáceres",
    "Carlos Sotelo sala",
    "Alejandro Tipula Diaz",
    "Barbara Urday Alata",
    "Cassandra Cárdenas Núñez",
    "Lucciana Kcana Olmedo"
    ]

    for i in range(len(student_names)):
        # print(f"{i}")
        student_name = student_names[i]
        apoderado = apoderados[i]
        num = contactos[i]
        print(f"""
{num}
Saludos cordiales {apoderado},
Soy José Vilca, miembro de ACM UCSP, y requiero el *correo electrónico* del estudiante *{student_name}* para concluir  su certificado de participación en el *Concurso Escolar de Programación 2023*

""")

if __name__ == "__main__":
    main()
