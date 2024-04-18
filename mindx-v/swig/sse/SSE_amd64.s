//+build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

TEXT ·_euclidean_distance_squared(SB), $0-32

    MOVQ len+0(FP), DI
    MOVQ a+8(FP), SI
    MOVQ b+16(FP), DX
    MOVQ result+24(FP), CX

    WORD $0x8949; BYTE $0xf8     // mov    r8, rdi
    LONG $0xfce08349             // and    r8, -4
	JE LBB0_1
    LONG $0xff508d4d             // lea    r10, [r8 - 1]
    WORD $0x894c; BYTE $0xd0     // mov    rax, r10
    LONG $0x02e8c148             // shr    rax, 2
    LONG $0x01488d44             // lea    r9d, [rax + 1]
    LONG $0x03e18341             // and    r9d, 3
    LONG $0x0cfa8349             // cmp    r10, 12
	JAE LBB0_12
    WORD $0x570f; BYTE $0xc0     // xorps    xmm0, xmm0
    WORD $0xc031                 // xor    eax, eax
    WORD $0x854d; BYTE $0xc9     // test    r9, r9
	JNE LBB0_5
	JMP LBB0_7
LBB0_1:
    WORD $0x570f; BYTE $0xc0     // xorps    xmm0, xmm0
	JMP LBB0_7
LBB0_12:
    LONG $0xff518d4d             // lea    r10, [r9 - 1]
    WORD $0x2949; BYTE $0xc2     // sub    r10, rax
    WORD $0x570f; BYTE $0xc9     // xorps    xmm1, xmm1
    WORD $0xc031                 // xor    eax, eax
LBB0_13:
    LONG $0x8614280f             // movaps    xmm2, oword [rsi + 4*rax]
    LONG $0x865c280f; BYTE $0x10 // movaps    xmm3, oword [rsi + 4*rax + 16]
    LONG $0x8664280f; BYTE $0x20 // movaps    xmm4, oword [rsi + 4*rax + 32]
    LONG $0x8644280f; BYTE $0x30 // movaps    xmm0, oword [rsi + 4*rax + 48]
    LONG $0x82145c0f             // subps    xmm2, oword [rdx + 4*rax]
    WORD $0x590f; BYTE $0xd2     // mulps    xmm2, xmm2
    WORD $0x580f; BYTE $0xd1     // addps    xmm2, xmm1
    LONG $0x825c5c0f; BYTE $0x10 // subps    xmm3, oword [rdx + 4*rax + 16]
    WORD $0x590f; BYTE $0xdb     // mulps    xmm3, xmm3
    WORD $0x580f; BYTE $0xda     // addps    xmm3, xmm2
    LONG $0x82645c0f; BYTE $0x20 // subps    xmm4, oword [rdx + 4*rax + 32]
    WORD $0x590f; BYTE $0xe4     // mulps    xmm4, xmm4
    WORD $0x580f; BYTE $0xe3     // addps    xmm4, xmm3
    LONG $0x82445c0f; BYTE $0x30 // subps    xmm0, oword [rdx + 4*rax + 48]
    WORD $0x590f; BYTE $0xc0     // mulps    xmm0, xmm0
    WORD $0x580f; BYTE $0xc4     // addps    xmm0, xmm4
    LONG $0x10c08348             // add    rax, 16
    WORD $0x280f; BYTE $0xc8     // movaps    xmm1, xmm0
    LONG $0x04c28349             // add    r10, 4
	JNE LBB0_13
    WORD $0x854d; BYTE $0xc9     // test    r9, r9
	JE LBB0_7
LBB0_5:
    LONG $0x02e0c148             // shl    rax, 2
    WORD $0xf749; BYTE $0xd9     // neg    r9
LBB0_6:
    LONG $0x060c280f             // movaps    xmm1, oword [rsi + rax]
    LONG $0x020c5c0f             // subps    xmm1, oword [rdx + rax]
    WORD $0x590f; BYTE $0xc9     // mulps    xmm1, xmm1
    WORD $0x580f; BYTE $0xc1     // addps    xmm0, xmm1
    LONG $0x10c08348             // add    rax, 16
    WORD $0xff49; BYTE $0xc1     // inc    r9
	JNE LBB0_6
LBB0_7:
    LONG $0xc8160ff3             // movshdup    xmm1, xmm0
    LONG $0xc8580ff3             // addss    xmm1, xmm0
    WORD $0x280f; BYTE $0xd0     // movaps    xmm2, xmm0
    LONG $0xd0150f66             // unpckhpd    xmm2, xmm0
    LONG $0xd1580ff3             // addss    xmm2, xmm1
    LONG $0xe7c0c60f             // shufps    xmm0, xmm0, 231
    LONG $0xc2580ff3             // addss    xmm0, xmm2
    LONG $0x01110ff3             // movss    dword [rcx], xmm0
    WORD $0x6349; BYTE $0xc0     // movsxd    rax, r8d
    WORD $0x3948; BYTE $0xf8     // cmp    rax, rdi
	JAE LBB0_11
    WORD $0x6349; BYTE $0xc0     // movsxd    rax, r8d
    WORD $0x634c; BYTE $0xc7     // movsxd    r8, edi
    WORD $0xf749; BYTE $0xd0     // not    r8
    LONG $0x03c88349             // or    r8, 3
    LONG $0x01c7f640             // test    dil, 1
	JE LBB0_10
    LONG $0x0c100ff3; BYTE $0x86 // movss    xmm1, dword [rsi + 4*rax]
    LONG $0x0c5c0ff3; BYTE $0x82 // subss    xmm1, dword [rdx + 4*rax]
    LONG $0xc9590ff3             // mulss    xmm1, xmm1
    LONG $0xc1580ff3             // addss    xmm0, xmm1
    LONG $0x01110ff3             // movss    dword [rcx], xmm0
    LONG $0x01c88348             // or    rax, 1
LBB0_10:
    WORD $0x0149; BYTE $0xf8     // add    r8, rdi
	JE LBB0_11
LBB0_14:
    LONG $0x0c100ff3; BYTE $0x86 // movss    xmm1, dword [rsi + 4*rax]
    LONG $0x0c5c0ff3; BYTE $0x82 // subss    xmm1, dword [rdx + 4*rax]
    LONG $0xc9590ff3             // mulss    xmm1, xmm1
    LONG $0xc8580ff3             // addss    xmm1, xmm0
    LONG $0x09110ff3             // movss    dword [rcx], xmm1
    LONG $0x44100ff3; WORD $0x0486 // movss    xmm0, dword [rsi + 4*rax + 4]
    LONG $0x445c0ff3; WORD $0x0482 // subss    xmm0, dword [rdx + 4*rax + 4]
    LONG $0xc0590ff3             // mulss    xmm0, xmm0
    LONG $0xc1580ff3             // addss    xmm0, xmm1
    LONG $0x01110ff3             // movss    dword [rcx], xmm0
    LONG $0x02c08348             // add    rax, 2
    WORD $0x3948; BYTE $0xc7     // cmp    rdi, rax
	JNE LBB0_14
LBB0_11:
    RET




DATA LCDATA1<>+0x000(SB)/8, $0x8000000080000000
DATA LCDATA1<>+0x008(SB)/8, $0x8000000080000000
GLOBL LCDATA1<>(SB), 8, $16

TEXT ·_manhattan_distance(SB), $0-32

    MOVQ len+0(FP), DI
    MOVQ a+8(FP), SI
    MOVQ b+16(FP), DX
    MOVQ result+24(FP), CX
    LEAQ LCDATA1<>(SB), BP

    WORD $0x8949; BYTE $0xf9     // mov    r9, rdi
    LONG $0xfce18349             // and    r9, -4
	JE LBB1_1
    LONG $0xff418d49             // lea    rax, [r9 - 1]
    LONG $0x02e8c148             // shr    rax, 2
    LONG $0x01408d44             // lea    r8d, [rax + 1]
    LONG $0x01e08341             // and    r8d, 1
    WORD $0x8548; BYTE $0xc0     // test    rax, rax
	JE LBB1_3
    LONG $0xff508d4d             // lea    r10, [r8 - 1]
    WORD $0x2949; BYTE $0xc2     // sub    r10, rax
    WORD $0x570f; BYTE $0xc0     // xorps    xmm0, xmm0
    WORD $0xc031                 // xor    eax, eax
LBB1_14:
    LONG $0x860c280f             // movaps    xmm1, oword [rsi + 4*rax]
    LONG $0x8654280f; BYTE $0x10 // movaps    xmm2, oword [rsi + 4*rax + 16]
    LONG $0x820c5c0f             // subps    xmm1, oword [rdx + 4*rax]
    WORD $0x590f; BYTE $0xc9     // mulps    xmm1, xmm1
    WORD $0x510f; BYTE $0xc9     // sqrtps    xmm1, xmm1
    WORD $0x580f; BYTE $0xc8     // addps    xmm1, xmm0
    LONG $0x82545c0f; BYTE $0x10 // subps    xmm2, oword [rdx + 4*rax + 16]
    WORD $0x590f; BYTE $0xd2     // mulps    xmm2, xmm2
    WORD $0x510f; BYTE $0xc2     // sqrtps    xmm0, xmm2
    WORD $0x580f; BYTE $0xc1     // addps    xmm0, xmm1
    LONG $0x08c08348             // add    rax, 8
    LONG $0x02c28349             // add    r10, 2
	JNE LBB1_14
    WORD $0x854d; BYTE $0xc0     // test    r8, r8
	JNE LBB1_5
	JMP LBB1_6
LBB1_1:
    WORD $0x570f; BYTE $0xc0     // xorps    xmm0, xmm0
	JMP LBB1_6
LBB1_3:
    WORD $0x570f; BYTE $0xc0     // xorps    xmm0, xmm0
    WORD $0xc031                 // xor    eax, eax
    WORD $0x854d; BYTE $0xc0     // test    r8, r8
	JE LBB1_6
LBB1_5:
    LONG $0x860c280f             // movaps    xmm1, oword [rsi + 4*rax]
    LONG $0x820c5c0f             // subps    xmm1, oword [rdx + 4*rax]
    WORD $0x590f; BYTE $0xc9     // mulps    xmm1, xmm1
    WORD $0x510f; BYTE $0xc9     // sqrtps    xmm1, xmm1
    WORD $0x580f; BYTE $0xc1     // addps    xmm0, xmm1
LBB1_6:
    LONG $0xc8160ff3             // movshdup    xmm1, xmm0
    LONG $0xc8580ff3             // addss    xmm1, xmm0
    WORD $0x280f; BYTE $0xd0     // movaps    xmm2, xmm0
    LONG $0xd0150f66             // unpckhpd    xmm2, xmm0
    LONG $0xd1580ff3             // addss    xmm2, xmm1
    LONG $0xe7c0c60f             // shufps    xmm0, xmm0, 231
    LONG $0xc2580ff3             // addss    xmm0, xmm2
    LONG $0x01110ff3             // movss    dword [rcx], xmm0
    WORD $0x6349; BYTE $0xc1     // movsxd    rax, r9d
    WORD $0x3948; BYTE $0xf8     // cmp    rax, rdi
	JAE LBB1_12
    WORD $0x6349; BYTE $0xc1     // movsxd    rax, r9d
    WORD $0x634c; BYTE $0xc7     // movsxd    r8, edi
    WORD $0xf749; BYTE $0xd0     // not    r8
    LONG $0x03c88349             // or    r8, 3
    LONG $0x01c7f640             // test    dil, 1
	JE LBB1_9
    LONG $0x0c100ff3; BYTE $0x86 // movss    xmm1, dword [rsi + 4*rax]
    LONG $0x0c5c0ff3; BYTE $0x82 // subss    xmm1, dword [rdx + 4*rax]
    LONG $0x0055280f             // movaps    xmm2, oword 0[rbp] /* [rip + LCPI1_0] */
    WORD $0x570f; BYTE $0xd1     // xorps    xmm2, xmm1
    WORD $0x570f; BYTE $0xdb     // xorps    xmm3, xmm3
    WORD $0x280f; BYTE $0xe1     // movaps    xmm4, xmm1
    LONG $0xe3c20ff3; BYTE $0x01 // cmpltss    xmm4, xmm3
    WORD $0x280f; BYTE $0xdc     // movaps    xmm3, xmm4
    WORD $0x550f; BYTE $0xd9     // andnps    xmm3, xmm1
    WORD $0x540f; BYTE $0xe2     // andps    xmm4, xmm2
    WORD $0x560f; BYTE $0xe3     // orps    xmm4, xmm3
    LONG $0xc4580ff3             // addss    xmm0, xmm4
    LONG $0x01110ff3             // movss    dword [rcx], xmm0
    LONG $0x01c88348             // or    rax, 1
LBB1_9:
    WORD $0x0149; BYTE $0xf8     // add    r8, rdi
	JE LBB1_12
    LONG $0x004d280f             // movaps    xmm1, oword 0[rbp] /* [rip + LCPI1_0] */
    WORD $0x570f; BYTE $0xd2     // xorps    xmm2, xmm2
LBB1_11:
    LONG $0x1c100ff3; BYTE $0x86 // movss    xmm3, dword [rsi + 4*rax]
    LONG $0x1c5c0ff3; BYTE $0x82 // subss    xmm3, dword [rdx + 4*rax]
    WORD $0x280f; BYTE $0xe3     // movaps    xmm4, xmm3
    LONG $0xe2c20ff3; BYTE $0x01 // cmpltss    xmm4, xmm2
    WORD $0x280f; BYTE $0xec     // movaps    xmm5, xmm4
    WORD $0x550f; BYTE $0xeb     // andnps    xmm5, xmm3
    WORD $0x570f; BYTE $0xd9     // xorps    xmm3, xmm1
    WORD $0x540f; BYTE $0xe3     // andps    xmm4, xmm3
    WORD $0x560f; BYTE $0xe5     // orps    xmm4, xmm5
    LONG $0xe0580ff3             // addss    xmm4, xmm0
    LONG $0x21110ff3             // movss    dword [rcx], xmm4
    LONG $0x5c100ff3; WORD $0x0486 // movss    xmm3, dword [rsi + 4*rax + 4]
    LONG $0x5c5c0ff3; WORD $0x0482 // subss    xmm3, dword [rdx + 4*rax + 4]
    WORD $0x280f; BYTE $0xc3     // movaps    xmm0, xmm3
    LONG $0xc2c20ff3; BYTE $0x01 // cmpltss    xmm0, xmm2
    WORD $0x280f; BYTE $0xe8     // movaps    xmm5, xmm0
    WORD $0x550f; BYTE $0xeb     // andnps    xmm5, xmm3
    WORD $0x570f; BYTE $0xd9     // xorps    xmm3, xmm1
    WORD $0x540f; BYTE $0xc3     // andps    xmm0, xmm3
    WORD $0x560f; BYTE $0xc5     // orps    xmm0, xmm5
    LONG $0xc4580ff3             // addss    xmm0, xmm4
    LONG $0x01110ff3             // movss    dword [rcx], xmm0
    LONG $0x02c08348             // add    rax, 2
    WORD $0x3948; BYTE $0xc7     // cmp    rdi, rax
	JNE LBB1_11
LBB1_12:
    RET




TEXT ·_cosine_similarity_dot_norm(SB), $0-40

    MOVQ len+0(FP), DI
    MOVQ a+8(FP), SI
    MOVQ b+16(FP), DX
    MOVQ dot+24(FP), CX
    MOVQ norm_squared+32(FP), R8

    WORD $0x8949; BYTE $0xfa     // mov    r10, rdi
    LONG $0xfce28349             // and    r10, -4
	JE LBB2_1
    LONG $0xff428d49             // lea    rax, [r10 - 1]
    LONG $0x02e8c148             // shr    rax, 2
    LONG $0x01488d44             // lea    r9d, [rax + 1]
    LONG $0x01e18341             // and    r9d, 1
    WORD $0x8548; BYTE $0xc0     // test    rax, rax
	JE LBB2_3
    LONG $0xff598d4d             // lea    r11, [r9 - 1]
    WORD $0x2949; BYTE $0xc3     // sub    r11, rax
    WORD $0x570f; BYTE $0xdb     // xorps    xmm3, xmm3
    WORD $0xc031                 // xor    eax, eax
    WORD $0x570f; BYTE $0xe4     // xorps    xmm4, xmm4
    WORD $0x570f; BYTE $0xc9     // xorps    xmm1, xmm1
LBB2_12:
    LONG $0x862c280f             // movaps    xmm5, oword [rsi + 4*rax]
    LONG $0x8654280f; BYTE $0x10 // movaps    xmm2, oword [rsi + 4*rax + 16]
    LONG $0x8234280f             // movaps    xmm6, oword [rdx + 4*rax]
    LONG $0x8244280f; BYTE $0x10 // movaps    xmm0, oword [rdx + 4*rax + 16]
    WORD $0x280f; BYTE $0xfd     // movaps    xmm7, xmm5
    WORD $0x590f; BYTE $0xfe     // mulps    xmm7, xmm6
    WORD $0x580f; BYTE $0xf9     // addps    xmm7, xmm1
    WORD $0x590f; BYTE $0xed     // mulps    xmm5, xmm5
    WORD $0x580f; BYTE $0xec     // addps    xmm5, xmm4
    WORD $0x590f; BYTE $0xf6     // mulps    xmm6, xmm6
    WORD $0x580f; BYTE $0xf3     // addps    xmm6, xmm3
    WORD $0x280f; BYTE $0xca     // movaps    xmm1, xmm2
    WORD $0x590f; BYTE $0xc8     // mulps    xmm1, xmm0
    WORD $0x580f; BYTE $0xcf     // addps    xmm1, xmm7
    WORD $0x590f; BYTE $0xd2     // mulps    xmm2, xmm2
    WORD $0x580f; BYTE $0xd5     // addps    xmm2, xmm5
    WORD $0x590f; BYTE $0xc0     // mulps    xmm0, xmm0
    WORD $0x580f; BYTE $0xc6     // addps    xmm0, xmm6
    LONG $0x08c08348             // add    rax, 8
    WORD $0x280f; BYTE $0xd8     // movaps    xmm3, xmm0
    WORD $0x280f; BYTE $0xe2     // movaps    xmm4, xmm2
    LONG $0x02c38349             // add    r11, 2
	JNE LBB2_12
    WORD $0x854d; BYTE $0xc9     // test    r9, r9
	JNE LBB2_5
	JMP LBB2_6
LBB2_1:
    WORD $0x570f; BYTE $0xc9     // xorps    xmm1, xmm1
    WORD $0x570f; BYTE $0xd2     // xorps    xmm2, xmm2
    WORD $0x570f; BYTE $0xc0     // xorps    xmm0, xmm0
	JMP LBB2_6
LBB2_3:
    WORD $0x570f; BYTE $0xc0     // xorps    xmm0, xmm0
    WORD $0xc031                 // xor    eax, eax
    WORD $0x570f; BYTE $0xd2     // xorps    xmm2, xmm2
    WORD $0x570f; BYTE $0xc9     // xorps    xmm1, xmm1
    WORD $0x854d; BYTE $0xc9     // test    r9, r9
	JE LBB2_6
LBB2_5:
    LONG $0x861c280f             // movaps    xmm3, oword [rsi + 4*rax]
    LONG $0x8224280f             // movaps    xmm4, oword [rdx + 4*rax]
    WORD $0x280f; BYTE $0xeb     // movaps    xmm5, xmm3
    WORD $0x590f; BYTE $0xdc     // mulps    xmm3, xmm4
    WORD $0x590f; BYTE $0xe4     // mulps    xmm4, xmm4
    WORD $0x580f; BYTE $0xc4     // addps    xmm0, xmm4
    WORD $0x590f; BYTE $0xed     // mulps    xmm5, xmm5
    WORD $0x580f; BYTE $0xd5     // addps    xmm2, xmm5
    WORD $0x580f; BYTE $0xcb     // addps    xmm1, xmm3
LBB2_6:
    LONG $0xd9160ff3             // movshdup    xmm3, xmm1
    LONG $0xd9580ff3             // addss    xmm3, xmm1
    WORD $0x280f; BYTE $0xe1     // movaps    xmm4, xmm1
    LONG $0xe1150f66             // unpckhpd    xmm4, xmm1
    LONG $0xe3580ff3             // addss    xmm4, xmm3
    LONG $0xe7c9c60f             // shufps    xmm1, xmm1, 231
    LONG $0xcc580ff3             // addss    xmm1, xmm4
    LONG $0x09110ff3             // movss    dword [rcx], xmm1
    WORD $0x280f; BYTE $0xda     // movaps    xmm3, xmm2
    LONG $0x213a0f66; WORD $0x1cd8 // insertps    xmm3, xmm0, 28
    WORD $0x280f; BYTE $0xe0     // movaps    xmm4, xmm0
    LONG $0x213a0f66; WORD $0x4ce2 // insertps    xmm4, xmm2, 76
    WORD $0x580f; BYTE $0xe3     // addps    xmm4, xmm3
    WORD $0x280f; BYTE $0xda     // movaps    xmm3, xmm2
    LONG $0xda150f66             // unpckhpd    xmm3, xmm2
    LONG $0x213a0f66; WORD $0x9cd8 // insertps    xmm3, xmm0, 156
    WORD $0x580f; BYTE $0xdc     // addps    xmm3, xmm4
    WORD $0x120f; BYTE $0xc0     // movhlps    xmm0, xmm0
    LONG $0x213a0f66; WORD $0xccc2 // insertps    xmm0, xmm2, 204
    WORD $0x580f; BYTE $0xc3     // addps    xmm0, xmm3
    WORD $0x6349; BYTE $0xc2     // movsxd    rax, r10d
    WORD $0x3948; BYTE $0xf8     // cmp    rax, rdi
	JAE LBB2_10
    WORD $0x6349; BYTE $0xc2     // movsxd    rax, r10d
    WORD $0x634c; BYTE $0xcf     // movsxd    r9, edi
    WORD $0xf749; BYTE $0xd1     // not    r9
    LONG $0x03c98349             // or    r9, 3
    LONG $0x01c7f640             // test    dil, 1
	JE LBB2_9
    LONG $0x14100ff3; BYTE $0x86 // movss    xmm2, dword [rsi + 4*rax]
    LONG $0x14590ff3; BYTE $0x82 // mulss    xmm2, dword [rdx + 4*rax]
    LONG $0xca580ff3             // addss    xmm1, xmm2
    LONG $0x09110ff3             // movss    dword [rcx], xmm1
    LONG $0x14100ff3; BYTE $0x86 // movss    xmm2, dword [rsi + 4*rax]
    LONG $0x213a0f66; WORD $0x8214; BYTE $0x10 // insertps    xmm2, dword [rdx + 4*rax], 16
    WORD $0x590f; BYTE $0xd2     // mulps    xmm2, xmm2
    WORD $0x580f; BYTE $0xc2     // addps    xmm0, xmm2
    LONG $0x01c88348             // or    rax, 1
LBB2_9:
    WORD $0x0149; BYTE $0xf9     // add    r9, rdi
	JE LBB2_10
LBB2_13:
    LONG $0x14100ff3; BYTE $0x86 // movss    xmm2, dword [rsi + 4*rax]
    LONG $0x14590ff3; BYTE $0x82 // mulss    xmm2, dword [rdx + 4*rax]
    LONG $0xd1580ff3             // addss    xmm2, xmm1
    LONG $0x11110ff3             // movss    dword [rcx], xmm2
    LONG $0x1c100ff3; BYTE $0x86 // movss    xmm3, dword [rsi + 4*rax]
    LONG $0x4c100ff3; WORD $0x0486 // movss    xmm1, dword [rsi + 4*rax + 4]
    LONG $0x213a0f66; WORD $0x821c; BYTE $0x10 // insertps    xmm3, dword [rdx + 4*rax], 16
    WORD $0x590f; BYTE $0xdb     // mulps    xmm3, xmm3
    WORD $0x580f; BYTE $0xd8     // addps    xmm3, xmm0
    LONG $0x4c590ff3; WORD $0x0482 // mulss    xmm1, dword [rdx + 4*rax + 4]
    LONG $0xca580ff3             // addss    xmm1, xmm2
    LONG $0x09110ff3             // movss    dword [rcx], xmm1
    LONG $0x44100ff3; WORD $0x0486 // movss    xmm0, dword [rsi + 4*rax + 4]
    QUAD $0x10048244213a0f66     // insertps    xmm0, dword [rdx + 4*rax + 4], 16
    WORD $0x590f; BYTE $0xc0     // mulps    xmm0, xmm0
    WORD $0x580f; BYTE $0xc3     // addps    xmm0, xmm3
    LONG $0x02c08348             // add    rax, 2
    WORD $0x3948; BYTE $0xc7     // cmp    rdi, rax
	JNE LBB2_13
LBB2_10:
    LONG $0xc8160ff3             // movshdup    xmm1, xmm0
    LONG $0xc8590ff3             // mulss    xmm1, xmm0
    LONG $0x110f41f3; BYTE $0x08 // movss    dword [r8], xmm1
    RET