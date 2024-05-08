# SMFilter

## Que es?

Este buscador simple de maquinas escrito en Go.
Se visualiza una tabla para encontrar maquinas con las caracteristicas que el usuario indique con flags.

## Motivos?

2 Motivos :
- Utilizar Go para aprender 
- Buscar mi propia solución de filtrado de maquinas en una terminal ,debido a la falta de exfiltros 

Ej: Dificultad "Easy" y pero no poder exfiltrar maquinas "Easy" en [infosecmachines](https://infosecmachines.io/).

Todas las Maquinas han sido realizadas por el profe S4vitar.

## Instalación

```bash
git clone https://github.com/Red-Clay/SMFilter.git
go build tool # Linux
go build tool.exe #Windows

```

## Filtros

- Si se quiere exfiltrar por el nombre del argumento se debe añadir al principio una exclamación "!"

| Flags                 | Descripción                                             |
|-----------------------|---------------------------------------------------------|
| `-max`                | Numero maximo de maquina que se viusalizaran            |
| `-n`                  | Buscar maquina por el nombre (sin Insensitive Case).    |
| `-p`                  | Buscar maquina por la plataforma .                      |           
| `-c`                  | Buscar maquina por la certificación (no admite varios). |
| `-d`                  | Buscar maquina por la dificultad .                      |
| `-o`                  | Buscar maquina por el sistema operativo.                |
| `-t`                  | Buscar maquina por la tecnica utilizada.                |
| `-h`                  | Imprimir el uso de la herramienta y listar las flags    |

#### Ejemplo completo
```bash
tool -p HackTheBox -d !Easy -o Linux  -c !OSCP -t Enum
```



> [!NOTE] 
> El simbolo "!" se puede utilizar escapandolo con "\\!"
> 
> Se puede desactivar la caracteristica en una shell con:
> ```bash
>"setopt NO_BANG_HIST" >> ~/.bashrc  # (event history) 
>source ~/.bashrc
>```

### Ejemplos Terminal
#### Linux
![[images/LinuxGolangTool.png]]
#### Windows
![[images/WindowsGolangTool.png]]

