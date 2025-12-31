// TESTE DE FUNÇÂO UTILZANDO ASSEMBLY
// NOTA: go tem seu proprio "ASSEMBLY" > LEGAL!!!!
// TEXT ·NomeDaFuncao(SB), $0-24
// $0: Tamanho do frame local (não estamos usando variáveis locais)
// 24: Tamanho total dos argumentos (2 ints de 8 bytes + 1 retorno de 8 bytes = 24)
TEXT ·SomarCoisas(SB), $0-24
MOVQ a+0(FP), AX    // Move o primeiro argumento (a) para o registrador AX
    MOVQ b+8(FP), CX    // Move o segundo argumento (b) para o registrador CX
    ADDQ CX, AX         // Soma CX em AX (AX = AX + CX)
    MOVQ AX, ret+16(FP) // Move o resultado para o espaço de retorno na pilha
    RET                 // Retorna
