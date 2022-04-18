package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

var x86Assembly = []string{
	"AAA",
	"AAD",
	"AAM",
	"AAS",
	"ADC",
	"ADD",
	"AND",
	"ARPL",
	"BOUND",
	"BSF",
	"BSR",
	"BSWAP",
	"BT",
	"BTC",
	"BTR",
	"BTS",
	"CALL",
	"CBW",
	"CDQ",
	"CDQE",
	"CLC",
	"CLD",
	"CLFLUSH",
	"CLI",
	"CLTS",
	"CMC",
	"CMOVA",
	"CMOVAE",
	"CMOVB",
	"CMOVBE",
	"CMOVC",
	"CMOVE",
	"CMOVG",
	"CMOVGE",
	"CMOVL",
	"CMOVLE",
	"CMOVNA",
	"CMOVNAE",
	"CMOVNB",
	"CMOVNBE",
	"CMOVNC",
	"CMOVNE",
	"CMOVNG",
	"CMOVNGE",
	"CMOVNL",
	"CMOVNLE",
	"CMOVNO",
	"CMOVNP",
	"CMOVNS",
	"CMOVNZ",
	"CMOVO",
	"CMOVP",
	"CMOVPE",
	"CMOVPO",
	"CMOVS",
	"CMOVZ",
	"CMP",
	"CMPS",
	"CMPSB",
	"CMPSD",
	"CMPSQ",
	"CMPSW",
	"CMPXCHG",
	"CMPXCHG16B",
	"CMPXCHG8B",
	"CPUID",
	"CQO",
	"CRC32",
	"CVTPD2PI",
	"CVTPI2PD",
	"CVTPI2PS",
	"CVTPS2PI",
	"CVTTPD2PI",
	"CVTTPS2PI",
	"CWD",
	"CWDE",
	"DAA",
	"DAS",
	"DEC",
	"DIV",
	"EMMS",
	"ENTER",
	"F2XM1",
	"FABS",
	"FADD",
	"FADDP",
	"FBLD",
	"FBSTP",
	"FCHS",
	"FCLEX",
	"FCMOVB",
	"FCMOVBE",
	"FCMOVE",
	"FCMOVNB",
	"FCMOVNBE",
	"FCMOVNE",
	"FCMOVNU",
	"FCMOVU",
	"FCOM",
	"FCOMI",
	"FCOMIP",
	"FCOMP",
	"FCOMPP",
	"FCOS",
	"FDECSTP",
	"FDIV",
	"FDIVP",
	"FDIVR",
	"FDIVRP",
	"FFREE",
	"FFREEP",
	"FIADD",
	"FICOM",
	"FICOMP",
	"FIDIV",
	"FIDIVR",
	"FILD",
	"FIMUL",
	"FINCSTP",
	"FINIT",
	"FIST",
	"FISTP",
	"FISTTP",
	"FISUB",
	"FISUBR",
	"FLD",
	"FLD1",
	"FLDCW",
	"FLDENV",
	"FLDL2E",
	"FLDL2T",
	"FLDLG2",
	"FLDLN2",
	"FLDPI",
	"FLDZ",
	"FMUL",
	"FMULP",
	"FNCLEX",
	"FNINIT",
	"FNOP",
	"FNSAVE",
	"FNSTCW",
	"FNSTENV",
	"FNSTSW",
	"FPATAN",
	"FPREM",
	"FPREM1",
	"FPTAN",
	"FRNDINT",
	"FRSTOR",
	"FSAVE",
	"FSCALE",
	"FSIN",
	"FSINCOS",
	"FSQRT",
	"FST",
	"FSTCW",
	"FSTENV",
	"FSTP",
	"FSTSW",
	"FSUB",
	"FSUBP",
	"FSUBR",
	"FSUBRP",
	"FTST",
	"FUCOM",
	"FUCOMI",
	"FUCOMIP",
	"FUCOMP",
	"FUCOMPP",
	"FWAIT",
	"FXAM",
	"FXCH",
	"FXRSTOR",
	"FXRSTOR64",
	"FXSAVE",
	"FXSAVE64",
	"FXTRACT",
	"FYL2X",
	"FYL2XP1",
	"HLT",
	"ICEBP",
	"IDIV",
	"IMUL",
	"IN",
	"INC",
	"INS",
	"INSB",
	"INSD",
	"INSW",
	"INT",
	"INTO",
	"INVD",
	"INVLPG",
	"IRET",
	"IRETD",
	"IRETQ",
	"JA",
	"JAE",
	"JB",
	"JBE",
	"JC",
	"JCXZ",
	"JE",
	"JECXZ",
	"JG",
	"JGE",
	"JL",
	"JLE",
	"JMP",
	"JNA",
	"JNAE",
	"JNB",
	"JNBE",
	"JNC",
	"JNE",
	"JNG",
	"JNGE",
	"JNL",
	"JNLE",
	"JNO",
	"JNP",
	"JNS",
	"JNZ",
	"JO",
	"JP",
	"JPE",
	"JPO",
	"JRCXZ",
	"JS",
	"JZ",
	"LAHF",
	"LAR",
	"LCALL",
	"LDS",
	"LEA",
	"LEAVE",
	"LES",
	"LFENCE",
	"LFS",
	"LGDT",
	"LGS",
	"LIDT",
	"LJMP",
	"LLDT",
	"LMSW",
	"LOCK",
	"LODS",
	"LODSB",
	"LODSD",
	"LODSQ",
	"LODSW",
	"LOOP",
	"LOOPE",
	"LOOPNE",
	"LRET",
	"LSL",
	"LSS",
	"LTR",
	"MASKMOVQ",
	"MFENCE",
	"MONITOR",
	"MOV",
	"MOVBE",
	"MOVDQ2Q",
	"MOVNTI",
	"MOVNTQ",
	"MOVQ2DQ",
	"MOVS",
	"MOVSB",
	"MOVSD",
	"MOVSQ",
	"MOVSW",
	"MOVSX",
	"MOVSXD",
	"MOVZX",
	"MUL",
	"MWAIT",
	"NEG",
	"NOP",
	"NOT",
	"OR",
	"OUT",
	"OUTS",
	"OUTSB",
	"OUTSD",
	"OUTSW",
	"PAUSE",
	"POP",
	"POPA",
	"POPAD",
	"POPF",
	"POPFD",
	"POPFQ",
	"PREFETCHNTA",
	"PREFETCHT0",
	"PREFETCHT1",
	"PREFETCHT2",
	"PSHUFW",
	"PUSH",
	"PUSHA",
	"PUSHAD",
	"PUSHF",
	"PUSHFD",
	"PUSHFQ",
	"RCL",
	"RCR",
	"RDMSR",
	"RDPMC",
	"RDTSC",
	"RDTSCP",
	"REP",
	"REPE",
	"RET",
	"ROL",
	"ROR",
	"RSM",
	"SAHF",
	"SAL",
	"SAR",
	"SBB",
	"SCAS",
	"SCASB",
	"SCASD",
	"SCASQ",
	"SCASW",
	"SETA",
	"SETAE",
	"SETB",
	"SETBE",
	"SETC",
	"SETE",
	"SETG",
	"SETGE",
	"SETL",
	"SETLE",
	"SETNA",
	"SETNAE",
	"SETNB",
	"SETNBE",
	"SETNC",
	"SETNE",
	"SETNG",
	"SETNGE",
	"SETNL",
	"SETNLE",
	"SETNO",
	"SETNP",
	"SETNS",
	"SETNZ",
	"SETO",
	"SETP",
	"SETPE",
	"SETPO",
	"SETS",
	"SETZ",
	"SFENCE",
	"SGDT",
	"SHL",
	"SHLD",
	"SHR",
	"SHRD",
	"SIDT",
	"SLDT",
	"SMSW",
	"STC",
	"STD",
	"STI",
	"STOS",
	"STOSB",
	"STOSD",
	"STOSQ",
	"STOSW",
	"STR",
	"SUB",
	"SWAPGS",
	"SYSCALL",
	"SYSENTER",
	"SYSEXIT",
	"SYSRET",
	"TEST",
	"UD0",
	"UD1",
	"UD2",
	"VERR",
	"VERW",
	"WAIT",
	"WBINVD",
	"WRMSR",
	"XADD",
	"XCHG",
	"XGETBV",
	"XLAT",
	"XLATB",
	"XOR",
	"XRSTOR",
	"XRSTOR64",
	"XRSTORS",
	"XRSTORS64",
	"XSAVE",
	"XSAVE64",
	"XSAVEC",
	"XSAVEC64",
	"XSAVES",
	"XSAVES64",
	"XSETBV",
}

var sse = []string{
	"ADDPS",
	"ADDSS",
	"ANDNPS",
	"ANDPS",
	"CMPPS",
	"CMPSS",
	"COMISS",
	"CVTSI2SS",
	"CVTSS2SI",
	"CVTTSS2SI",
	"DIVPS",
	"DIVSS",
	"LDMXCSR",
	"MAXPS",
	"MAXSS",
	"MINPS",
	"MINSS",
	"MOVAPS",
	"MOVHLPS",
	"MOVHPS",
	"MOVLHPS",
	"MOVLPS",
	"MOVMSKPS",
	"MOVNTPS",
	"MOVNTSD",
	"MOVNTSS",
	"MOVSS",
	"MOVUPS",
	"MULPS",
	"MULSS",
	"ORPS",
	"PAVGB",
	"PAVGW",
	"PEXTRW",
	"PINSRW",
	"PMAXSW",
	"PMAXUB",
	"PMINSW",
	"PMINUB",
	"PMOVMSKB",
	"PMULHUW",
	"PSADBW",
	"RCPPS",
	"RCPSS",
	"RSQRTPS",
	"RSQRTSS",
	"SHUFPS",
	"SQRTPS",
	"SQRTSS",
	"STMXCSR",
	"SUBPS",
	"SUBSS",
	"UCOMISS",
	"UNPCKHPS",
	"UNPCKLPS",
	"XORPS",
}
var sse2 = []string{
	"ADDPD",
	"ADDSD",
	"ANDNPD",
	"ANDPD",
	"CMPPD",
	"CMPSD_XMM",
	"COMISD",
	"CVTDQ2PD",
	"CVTDQ2PS",
	"CVTPD2DQ",
	"CVTPD2PS",
	"CVTPS2DQ",
	"CVTPS2PD",
	"CVTSD2SI",
	"CVTSD2SS",
	"CVTSI2SD",
	"CVTSS2SD",
	"CVTTPD2DQ",
	"CVTTPS2DQ",
	"CVTTSD2SI",
	"DIVPD",
	"DIVSD",
	"MASKMOVDQU",
	"MAXPD",
	"MAXSD",
	"MINPD",
	"MINSD",
	"MOVAPD",
	"MOVD",
	"MOVDQA",
	"MOVDQU",
	"MOVHPD",
	"MOVLPD",
	"MOVMSKPD",
	"MOVNTDQ",
	"MOVNTPD",
	"MOVQ",
	"MOVSD_XMM",
	"MOVUPD",
	"MULPD",
	"MULSD",
	"ORPD",
	"PACKSSDW",
	"PACKSSWB",
	"PACKUSWB",
	"PADDB",
	"PADDD",
	"PADDQ",
	"PADDSB",
	"PADDSW",
	"PADDUSB",
	"PADDUSW",
	"PADDW",
	"PAND",
	"PANDN",
	"PAVGB",
	"PAVGW",
	"PCMPEQB",
	"PCMPEQD",
	"PCMPEQW",
	"PCMPGTB",
	"PCMPGTD",
	"PCMPGTW",
	"PEXTRW",
	"PINSRW",
	"PMADDWD",
	"PMAXSW",
	"PMAXUB",
	"PMINSW",
	"PMINUB",
	"PMOVMSKB",
	"PMULHUW",
	"PMULHW",
	"PMULLW",
	"PMULUDQ",
	"POR",
	"PSADBW",
	"PSHUFD",
	"PSHUFHW",
	"PSHUFLW",
	"PSLLD",
	"PSLLDQ",
	"PSLLQ",
	"PSLLW",
	"PSRAD",
	"PSRAW",
	"PSRLD",
	"PSRLDQ",
	"PSRLQ",
	"PSRLW",
	"PSUBB",
	"PSUBD",
	"PSUBQ",
	"PSUBSB",
	"PSUBSW",
	"PSUBUSB",
	"PSUBUSW",
	"PSUBW",
	"PUNPCKHBW",
	"PUNPCKHDQ",
	"PUNPCKHQDQ",
	"PUNPCKHWD",
	"PUNPCKLBW",
	"PUNPCKLDQ",
	"PUNPCKLQDQ",
	"PUNPCKLWD",
	"PXOR",
	"SHUFPD",
	"SQRTPD",
	"SQRTSD",
	"SUBPD",
	"SUBSD",
	"UCOMISD",
	"UNPCKHPD",
	"UNPCKLPD",
	"XORPD",
}

var x86SixyFourAssembly = append(append(x86Assembly, sse...), sse2...)

// https://github.com/golang/arch/blob/master/x86/x86.csv
var cmpxchg16b = []string{
	"CMPXCHG16B",
}
var lahf = []string{
	"LAHF",
}
var sahf = []string{
	"SAHF",
}
var popcnt = []string{
	"POPCNT",
}
var sse3 = []string{
	"ADDSUBPD",
	"ADDSUBPS",
	"HADDPD",
	"HADDPS",
	"HSUBPD",
	"HSUBPS",
	"LDDQU",
	"MOVDDUP",
	"MOVSHDUP",
	"MOVSLDUP",
}
var sse41 = []string{
	"BLENDPD",
	"BLENDPS",
	"BLENDVPD",
	"BLENDVPS",
	"DPPD",
	"DPPS",
	"EXTRACTPS",
	"INSERTPS",
	"MOVNTDQA",
	"MPSADBW",
	"PACKUSDW",
	"PBLENDVB",
	"PBLENDW",
	"PCMPEQQ",
	"PEXTRB",
	"PEXTRD",
	"PEXTRQ",
	"PEXTRW",
	"PHMINPOSUW",
	"PINSRB",
	"PINSRD",
	"PINSRQ",
	"PMAXSB",
	"PMAXSD",
	"PMAXUD",
	"PMAXUW",
	"PMINSB",
	"PMINSD",
	"PMINUD",
	"PMINUW",
	"PMOVSXBD",
	"PMOVSXBQ",
	"PMOVSXBW",
	"PMOVSXDQ",
	"PMOVSXWD",
	"PMOVSXWQ",
	"PMOVZXBD",
	"PMOVZXBQ",
	"PMOVZXBW",
	"PMOVZXDQ",
	"PMOVZXWD",
	"PMOVZXWQ",
	"PMULDQ",
	"PMULLD",
	"PTEST",
	"ROUNDPD",
	"ROUNDPS",
	"ROUNDSD",
	"ROUNDSS",
}
var sse42 = []string{
	"PCMPESTRI",
	"PCMPESTRM",
	"PCMPGTQ",
	"PCMPISTRI",
	"PCMPISTRM",
}

var ssse3 = []string{
	"PABSB",
	"PABSD",
	"PABSW",
	"PALIGNR",
	"PHADDD",
	"PHADDSW",
	"PHADDW",
	"PHSUBD",
	"PHSUBSW",
	"PHSUBW",
	"PMADDUBSW",
	"PMULHRSW",
	"PSHUFB",
	"PSIGNB",
	"PSIGND",
	"PSIGNW",
}
var v2Assembly = append(append(append(append(append(append(append(cmpxchg16b, lahf...), sahf...), popcnt...), sse3...), sse41...), sse42...), ssse3...)

var avx1 = []string{
	"VADDPD",
	"VADDPS",
	"VADDSD",
	"VADDSS",
	"VADDSUBPD",
	"VADDSUBPS",
	"VANDNPD",
	"VANDNPS",
	"VANDPD",
	"VANDPS",
	"VBLENDPD",
	"VBLENDPS",
	"VBLENDVPD",
	"VBLENDVPS",
	"VBROADCASTF128",
	"VBROADCASTSD",
	"VBROADCASTSS",
	"VCMPPD",
	"VCMPPS",
	"VCMPSD",
	"VCMPSS",
	"VCOMISD",
	"VCOMISS",
	"VCVTDQ2PD",
	"VCVTDQ2PS",
	"VCVTPD2DQ",
	"VCVTPD2PS",
	"VCVTPS2DQ",
	"VCVTPS2PD",
	"VCVTSD2SI",
	"VCVTSD2SS",
	"VCVTSI2SD",
	"VCVTSI2SS",
	"VCVTSS2SD",
	"VCVTSS2SI",
	"VCVTTPD2DQ",
	"VCVTTPS2DQ",
	"VCVTTSD2SI",
	"VCVTTSS2SI",
	"VDIVPD",
	"VDIVPS",
	"VDIVSD",
	"VDIVSS",
	"VDPPD",
	"VDPPS",
	"VEXTRACTF128",
	"VEXTRACTPS",
	"VHADDPD",
	"VHADDPS",
	"VHSUBPD",
	"VHSUBPS",
	"VINSERTF128",
	"VINSERTPS",
	"VLDDQU",
	"VLDMXCSR",
	"VMASKMOVDQU",
	"VMASKMOVPD",
	"VMASKMOVPS",
	"VMAXPD",
	"VMAXPS",
	"VMAXSD",
	"VMAXSS",
	"VMINPD",
	"VMINPS",
	"VMINSD",
	"VMINSS",
	"VMOVAPD",
	"VMOVAPS",
	"VMOVD",
	"VMOVDDUP",
	"VMOVDQA",
	"VMOVDQU",
	"VMOVHLPS",
	"VMOVHPD",
	"VMOVHPS",
	"VMOVLHPS",
	"VMOVLPD",
	"VMOVLPS",
	"VMOVMSKPD",
	"VMOVMSKPS",
	"VMOVNTDQ",
	"VMOVNTDQA",
	"VMOVNTPD",
	"VMOVNTPS",
	"VMOVQ",
	"VMOVSD",
	"VMOVSHDUP",
	"VMOVSLDUP",
	"VMOVSS",
	"VMOVUPD",
	"VMOVUPS",
	"VMPSADBW",
	"VMULPD",
	"VMULPS",
	"VMULSD",
	"VMULSS",
	"VORPD",
	"VORPS",
	"VPABSB",
	"VPABSD",
	"VPABSW",
	"VPACKSSDW",
	"VPACKSSWB",
	"VPACKUSDW",
	"VPACKUSWB",
	"VPADDB",
	"VPADDD",
	"VPADDQ",
	"VPADDSB",
	"VPADDSW",
	"VPADDUSB",
	"VPADDUSW",
	"VPADDW",
	"VPALIGNR",
	"VPAND",
	"VPANDN",
	"VPAVGB",
	"VPAVGW",
	"VPBLENDVB",
	"VPBLENDW",
	"VPCMPEQB",
	"VPCMPEQD",
	"VPCMPEQQ",
	"VPCMPEQW",
	"VPCMPESTRI",
	"VPCMPESTRM",
	"VPCMPGTB",
	"VPCMPGTD",
	"VPCMPGTQ",
	"VPCMPGTW",
	"VPCMPISTRI",
	"VPCMPISTRM",
	"VPERM2F128",
	"VPERMILPD",
	"VPERMILPS",
	"VPEXTRB",
	"VPEXTRD",
	"VPEXTRQ",
	"VPEXTRW",
	"VPHADDD",
	"VPHADDSW",
	"VPHADDW",
	"VPHMINPOSUW",
	"VPHSUBD",
	"VPHSUBSW",
	"VPHSUBW",
	"VPINSRB",
	"VPINSRD",
	"VPINSRQ",
	"VPINSRW",
	"VPMADDUBSW",
	"VPMADDWD",
	"VPMAXSB",
	"VPMAXSD",
	"VPMAXSW",
	"VPMAXUB",
	"VPMAXUD",
	"VPMAXUW",
	"VPMINSB",
	"VPMINSD",
	"VPMINSW",
	"VPMINUB",
	"VPMINUD",
	"VPMINUW",
	"VPMOVMSKB",
	"VPMOVSXBD",
	"VPMOVSXBQ",
	"VPMOVSXBW",
	"VPMOVSXDQ",
	"VPMOVSXWD",
	"VPMOVSXWQ",
	"VPMOVZXBD",
	"VPMOVZXBQ",
	"VPMOVZXBW",
	"VPMOVZXDQ",
	"VPMOVZXWD",
	"VPMOVZXWQ",
	"VPMULDQ",
	"VPMULHRSW",
	"VPMULHUW",
	"VPMULHW",
	"VPMULLD",
	"VPMULLW",
	"VPMULUDQ",
	"VPOR",
	"VPSADBW",
	"VPSHUFB",
	"VPSHUFD",
	"VPSHUFHW",
	"VPSHUFLW",
	"VPSIGNB",
	"VPSIGND",
	"VPSIGNW",
	"VPSLLD",
	"VPSLLDQ",
	"VPSLLQ",
	"VPSLLW",
	"VPSRAD",
	"VPSRAW",
	"VPSRLD",
	"VPSRLDQ",
	"VPSRLQ",
	"VPSRLW",
	"VPSUBB",
	"VPSUBD",
	"VPSUBQ",
	"VPSUBSB",
	"VPSUBSW",
	"VPSUBUSB",
	"VPSUBUSW",
	"VPSUBW",
	"VPTEST",
	"VPUNPCKHBW",
	"VPUNPCKHDQ",
	"VPUNPCKHQDQ",
	"VPUNPCKHWD",
	"VPUNPCKLBW",
	"VPUNPCKLDQ",
	"VPUNPCKLQDQ",
	"VPUNPCKLWD",
	"VPXOR",
	"VRCPPS",
	"VRCPSS",
	"VROUNDPD",
	"VROUNDPS",
	"VROUNDSD",
	"VROUNDSS",
	"VRSQRTPS",
	"VRSQRTSS",
	"VSHUFPD",
	"VSHUFPS",
	"VSQRTPD",
	"VSQRTPS",
	"VSQRTSD",
	"VSQRTSS",
	"VSTMXCSR",
	"VSUBPD",
	"VSUBPS",
	"VSUBSD",
	"VSUBSS",
	"VTESTPD",
	"VTESTPS",
	"VUCOMISD",
	"VUCOMISS",
	"VUNPCKHPD",
	"VUNPCKHPS",
	"VUNPCKLPD",
	"VUNPCKLPS",
	"VXORPD",
	"VXORPS",
	"VZEROALL",
	"VZEROUPPER",
}
var avx2 = []string{
	"VBROADCASTI128",
	"VBROADCASTSD",
	"VBROADCASTSS",
	"VEXTRACTI128",
	"VGATHERDPD",
	"VGATHERDPS",
	"VGATHERQPD",
	"VGATHERQPS",
	"VINSERTI128",
	"VMOVNTDQA",
	"VMPSADBW",
	"VPABSB",
	"VPABSD",
	"VPABSW",
	"VPACKSSDW",
	"VPACKSSWB",
	"VPACKUSDW",
	"VPACKUSWB",
	"VPADDB",
	"VPADDD",
	"VPADDQ",
	"VPADDSB",
	"VPADDSW",
	"VPADDUSB",
	"VPADDUSW",
	"VPADDW",
	"VPALIGNR",
	"VPAND",
	"VPANDN",
	"VPAVGB",
	"VPAVGW",
	"VPBLENDD",
	"VPBLENDVB",
	"VPBLENDW",
	"VPBROADCASTB",
	"VPBROADCASTD",
	"VPBROADCASTQ",
	"VPBROADCASTW",
	"VPCMPEQD",
	"VPCMPEQQ",
	"VPCMPEQW",
	"VPCMPGTB",
	"VPCMPGTD",
	"VPCMPGTQ",
	"VPCMPGTW",
	"VPERM2I128",
	"VPERMD",
	"VPERMPD",
	"VPERMPS",
	"VPERMQ",
	"VPGATHERDD",
	"VPGATHERDQ",
	"VPGATHERQD",
	"VPGATHERQQ",
	"VPHADDD",
	"VPHADDSW",
	"VPHADDW",
	"VPHSUBD",
	"VPHSUBSW",
	"VPHSUBW",
	"VPMADDUBSW",
	"VPMADDWD",
	"VPMASKMOVD",
	"VPMASKMOVQ",
	"VPMAXSB",
	"VPMAXSD",
	"VPMAXSW",
	"VPMAXUB",
	"VPMAXUD",
	"VPMAXUW",
	"VPMINSB",
	"VPMINSD",
	"VPMINSW",
	"VPMINUB",
	"VPMINUD",
	"VPMINUW",
	"VPMOVMSKB",
	"VPMOVSXBD",
	"VPMOVSXBQ",
	"VPMOVSXBW",
	"VPMOVSXWD",
	"VPMOVZXBD",
	"VPMOVZXBQ",
	"VPMOVZXBW",
	"VPMOVZXWD",
	"VPMULDQ",
	"VPMULHRSW",
	"VPMULHUW",
	"VPMULHW",
	"VPMULLD",
	"VPMULLW",
	"VPMULUDQ",
	"VPOR",
	"VPSADBW",
	"VPSHUFB",
	"VPSHUFD",
	"VPSHUFHW",
	"VPSHUFLW",
	"VPSIGNB",
	"VPSIGND",
	"VPSIGNW",
	"VPSLLDQ",
	"VPSLLVD",
	"VPSLLVQ",
	"VPSLLW",
	"VPSRAD",
	"VPSRAVD",
	"VPSRAW",
	"VPSRLDQ",
	"VPSRLVD",
	"VPSRLVQ",
	"VPSRLW",
	"VPSUBB",
	"VPSUBD",
	"VPSUBQ",
	"VPSUBSB",
	"VPSUBSW",
	"VPSUBUSB",
	"VPSUBUSW",
	"VPSUBW",
	"VPUNPCKHBW",
	"VPUNPCKHDQ",
	"VPUNPCKHQDQ",
	"VPUNPCKHWD",
	"VPUNPCKLBW",
	"VPUNPCKLDQ",
	"VPUNPCKLQDQ",
	"VPUNPCKLWD",
	"VPXOR",
}
var bmi1 = []string{
	"ANDN",
	"BEXTR",
	"BLSI",
	"BLSMSK",
	"BLSR",
	"TZCNT",
}
var bmi2 = []string{
	"BZHI",
	"MULX",
	"PDEP",
	"PEXT",
	"RORX",
	"SARX",
	"SHLX",
	"SHRX",
}
var f16c = []string{
	"VCVTPH2PS",
	"VCVTPS2PH",
}
var fma = []string{
	"VFMADD132PD",
	"VFMADD132PS",
	"VFMADD132SD",
	"VFMADD132SS",
	"VFMADD213PD",
	"VFMADD213PS",
	"VFMADD213SD",
	"VFMADD213SS",
	"VFMADD231PD",
	"VFMADD231PS",
	"VFMADD231SD",
	"VFMADD231SS",
	"VFMADDSUB132PD",
	"VFMADDSUB132PS",
	"VFMADDSUB213PD",
	"VFMADDSUB213PS",
	"VFMADDSUB231PD",
	"VFMADDSUB231PS",
	"VFMSUB132PD",
	"VFMSUB132PS",
	"VFMSUB132SD",
	"VFMSUB132SS",
	"VFMSUB213PD",
	"VFMSUB213PS",
	"VFMSUB213SD",
	"VFMSUB213SS",
	"VFMSUB231PD",
	"VFMSUB231PS",
	"VFMSUB231SD",
	"VFMSUB231SS",
	"VFMSUBADD132PD",
	"VFMSUBADD132PS",
	"VFMSUBADD213PD",
	"VFMSUBADD213PS",
	"VFMSUBADD231PD",
	"VFMSUBADD231PS",
	"VFNMADD132PD",
	"VFNMADD132PS",
	"VFNMADD132SD",
	"VFNMADD132SS",
	"VFNMADD213PD",
	"VFNMADD213PS",
	"VFNMADD213SD",
	"VFNMADD213SS",
	"VFNMADD231PD",
	"VFNMADD231PS",
	"VFNMADD231SD",
	"VFNMADD231SS",
	"VFNMSUB132PD",
	"VFNMSUB132PS",
	"VFNMSUB132SD",
	"VFNMSUB132SS",
	"VFNMSUB213PD",
	"VFNMSUB213PS",
	"VFNMSUB213SD",
	"VFNMSUB213SS",
	"VFNMSUB231PD",
	"VFNMSUB231PS",
	"VFNMSUB231SD",
	"VFNMSUB231SS",
}
var lzcnt = []string{
	"LZCNT",
}

// var movbe = []string{}
var osxsave = []string{
	"OSXSAVE",
}
var v3Assembly = append(append(append(append(append(append(append(avx1, avx2...), bmi1...), bmi2...), f16c...), fma...), lzcnt...), osxsave...)

// https://raw.githubusercontent.com/intel-go/avx512counters/master/avx512_core_i9_7900x.csv
var avx512Assembly = []string{
	"KADDB",
	"KADDD",
	"KADDQ",
	"KADDW",
	"KANDB",
	"KANDD",
	"KANDNB",
	"KANDND",
	"KANDNQ",
	"KANDNW",
	"KANDQ",
	"KANDW",
	"KMOVB",
	"KMOVD",
	"KMOVQ",
	"KMOVW",
	"KNOTB",
	"KNOTD",
	"KNOTQ",
	"KNOTW",
	"KORB",
	"KORD",
	"KORQ",
	"KORTESTB",
	"KORTESTD",
	"KORTESTQ",
	"KORTESTW",
	"KORW",
	"KSHIFTLB",
	"KSHIFTLD",
	"KSHIFTLQ",
	"KSHIFTLW",
	"KSHIFTRB",
	"KSHIFTRD",
	"KSHIFTRQ",
	"KSHIFTRW",
	"KTESTB",
	"KTESTD",
	"KTESTQ",
	"KTESTW",
	"KUNPCKBW",
	"KUNPCKDQ",
	"KUNPCKWD",
	"KXNORB",
	"KXNORD",
	"KXNORQ",
	"KXNORW",
	"KXORB",
	"KXORD",
	"KXORQ",
	"KXORW",
	"VADDPD",
	"VADDPS",
	"VADDSD",
	"VADDSS",
	"VALIGND",
	"VALIGNQ",
	"VANDNPD",
	"VANDNPS",
	"VANDPD",
	"VANDPS",
	"VBLENDMPD",
	"VBLENDMPS",
	"VBROADCASTF32X2",
	"VBROADCASTF32X4",
	"VBROADCASTF32X8",
	"VBROADCASTF64X2",
	"VBROADCASTF64X4",
	"VBROADCASTI32X2",
	"VBROADCASTI32X4",
	"VBROADCASTI32X8",
	"VBROADCASTI64X2",
	"VBROADCASTI64X4",
	"VBROADCASTSD",
	"VBROADCASTSS",
	"VCMPPD",
	"VCMPPS",
	"VCMPSD",
	"VCMPSS",
	"VCOMISD",
	"VCOMISS",
	"VCOMPRESSPD",
	"VCOMPRESSPS",
	"VCVTDQ2PD",
	"VCVTDQ2PS",
	"VCVTPD2DQ",
	"VCVTPD2DQX",
	"VCVTPD2DQY",
	"VCVTPD2PS",
	"VCVTPD2PSX",
	"VCVTPD2PSY",
	"VCVTPD2QQ",
	"VCVTPD2UDQ",
	"VCVTPD2UDQX",
	"VCVTPD2UDQY",
	"VCVTPD2UQQ",
	"VCVTPH2PS",
	"VCVTPS2DQ",
	"VCVTPS2PD",
	"VCVTPS2PH",
	"VCVTPS2QQ",
	"VCVTPS2UDQ",
	"VCVTPS2UQQ",
	"VCVTQQ2PD",
	"VCVTQQ2PS",
	"VCVTQQ2PSX",
	"VCVTQQ2PSY",
	"VCVTSD2SI",
	"VCVTSD2SS",
	"VCVTSD2USIL",
	"VCVTSD2USIQ",
	"VCVTSI2SDL",
	"VCVTSI2SDQ",
	"VCVTSI2SSL",
	"VCVTSS2SD",
	"VCVTSS2SI",
	"VCVTSS2USIL",
	"VCVTSS2USIQ",
	"VCVTTPD2DQ",
	"VCVTTPD2DQX",
	"VCVTTPD2DQY",
	"VCVTTPD2QQ",
	"VCVTTPD2UDQ",
	"VCVTTPD2UDQX",
	"VCVTTPD2UDQY",
	"VCVTTPD2UQQ",
	"VCVTTPS2DQ",
	"VCVTTPS2QQ",
	"VCVTTPS2UDQ",
	"VCVTTPS2UQQ",
	"VCVTTSD2SI",
	"VCVTTSD2USIL",
	"VCVTTSD2USIQ",
	"VCVTTSS2SI",
	"VCVTTSS2SIQ",
	"VCVTTSS2USIL",
	"VCVTTSS2USIQ",
	"VCVTUDQ2PD",
	"VCVTUDQ2PS",
	"VCVTUQQ2PD",
	"VCVTUQQ2PS",
	"VCVTUQQ2PSX",
	"VCVTUQQ2PSY",
	"VCVTUSI2SDL",
	"VCVTUSI2SDQ",
	"VCVTUSI2SSL",
	"VCVTUSI2SSQ",
	"VDBPSADBW",
	"VDIVPD",
	"VDIVPS",
	"VDIVSD",
	"VDIVSS",
	"VEXPANDPD",
	"VEXPANDPS",
	"VEXTRACTF32X4",
	"VEXTRACTF32X8",
	"VEXTRACTF64X2",
	"VEXTRACTF64X4",
	"VEXTRACTI32X4",
	"VEXTRACTI32X8",
	"VEXTRACTI64X2",
	"VEXTRACTI64X4",
	"VFIXUPIMMPD",
	"VFIXUPIMMPS",
	"VFIXUPIMMSD",
	"VFIXUPIMMSS",
	"VFMADD132PD",
	"VFMADD132PS",
	"VFMADD132SD",
	"VFMADD132SS",
	"VFMADD213PD",
	"VFMADD213PS",
	"VFMADD213SD",
	"VFMADD213SS",
	"VFMADD231PD",
	"VFMADD231PS",
	"VFMADD231SD",
	"VFMADD231SS",
	"VFMADDSUB132PD",
	"VFMADDSUB132PS",
	"VFMADDSUB213PD",
	"VFMADDSUB213PS",
	"VFMADDSUB231PD",
	"VFMADDSUB231PS",
	"VFMSUB132PD",
	"VFMSUB132PS",
	"VFMSUB132SD",
	"VFMSUB132SS",
	"VFMSUB213PD",
	"VFMSUB213PS",
	"VFMSUB213SD",
	"VFMSUB213SS",
	"VFMSUB231PD",
	"VFMSUB231PS",
	"VFMSUB231SD",
	"VFMSUB231SS",
	"VFMSUBADD132PD",
	"VFMSUBADD132PS",
	"VFMSUBADD213PD",
	"VFMSUBADD213PS",
	"VFMSUBADD231PD",
	"VFMSUBADD231PS",
	"VFNMADD132PD",
	"VFNMADD132PS",
	"VFNMADD132SD",
	"VFNMADD132SS",
	"VFNMADD213PD",
	"VFNMADD213PS",
	"VFNMADD213SD",
	"VFNMADD213SS",
	"VFNMADD231PD",
	"VFNMADD231PS",
	"VFNMADD231SD",
	"VFNMADD231SS",
	"VFNMSUB132PD",
	"VFNMSUB132PS",
	"VFNMSUB132SD",
	"VFNMSUB132SS",
	"VFNMSUB213PD",
	"VFNMSUB213PS",
	"VFNMSUB213SD",
	"VFNMSUB213SS",
	"VFNMSUB231PD",
	"VFNMSUB231PS",
	"VFNMSUB231SD",
	"VFNMSUB231SS",
	"VFPCLASSPDX",
	"VFPCLASSPDY",
	"VFPCLASSPDZ",
	"VFPCLASSPSX",
	"VFPCLASSPSY",
	"VFPCLASSPSZ",
	"VFPCLASSSD",
	"VFPCLASSSS",
	"VGATHERDPD",
	"VGATHERDPS",
	"VGATHERQPD",
	"VGATHERQPS",
	"VGETEXPPD",
	"VGETEXPPS",
	"VGETEXPSD",
	"VGETEXPSS",
	"VGETMANTPD",
	"VGETMANTPS",
	"VGETMANTSD",
	"VGETMANTSS",
	"VINSERTF32X4",
	"VINSERTF32X8",
	"VINSERTF64X2",
	"VINSERTF64X4",
	"VINSERTI32X4",
	"VINSERTI32X8",
	"VINSERTI64X2",
	"VINSERTI64X4",
	"VMAXPD",
	"VMAXPS",
	"VMAXSD",
	"VMAXSS",
	"VMINPD",
	"VMINPS",
	"VMINSD",
	"VMINSS",
	"VMOVAPD",
	"VMOVAPS",
	"VMOVDDUP",
	"VMOVDQA32",
	"VMOVDQA64",
	"VMOVDQU16",
	"VMOVDQU32",
	"VMOVDQU64",
	"VMOVDQU8",
	"VMOVHPS",
	"VMOVLHPS",
	"VMOVNTDQ",
	"VMOVNTDQA",
	"VMOVNTPD",
	"VMOVNTPS",
	"VMOVSD",
	"VMOVSHDUP",
	"VMOVSLDUP",
	"VMOVSS",
	"VMOVUPD",
	"VMOVUPS",
	"VMULPD",
	"VMULPS",
	"VMULSD",
	"VMULSS",
	"VORPD",
	"VORPS",
	"VPABSB",
	"VPABSD",
	"VPABSQ",
	"VPABSW",
	"VPACKSSDW",
	"VPACKSSWB",
	"VPACKUSDW",
	"VPACKUSWB",
	"VPADDB",
	"VPADDD",
	"VPADDQ",
	"VPADDSB",
	"VPADDSW",
	"VPADDUSB",
	"VPADDUSW",
	"VPADDW",
	"VPALIGNR",
	"VPANDD",
	"VPANDND",
	"VPANDNQ",
	"VPANDQ",
	"VPAVGB",
	"VPAVGW",
	"VPBLENDMB",
	"VPBLENDMD",
	"VPBLENDMQ",
	"VPBLENDMW",
	"VPBROADCASTB",
	"VPBROADCASTD",
	"VPBROADCASTMB2Q",
	"VPBROADCASTMW2D",
	"VPBROADCASTQ",
	"VPBROADCASTW",
	"VPCMPB",
	"VPCMPD",
	"VPCMPEQB",
	"VPCMPEQD",
	"VPCMPEQQ",
	"VPCMPEQW",
	"VPCMPGTB",
	"VPCMPGTD",
	"VPCMPGTQ",
	"VPCMPGTW",
	"VPCMPQ",
	"VPCMPUB",
	"VPCMPUD",
	"VPCMPUQ",
	"VPCMPUW",
	"VPCMPW",
	"VPCOMPRESSD",
	"VPCOMPRESSQ",
	"VPCONFLICTD",
	"VPCONFLICTQ",
	"VPERMD",
	"VPERMI2D",
	"VPERMI2PD",
	"VPERMI2PS",
	"VPERMI2Q",
	"VPERMI2W",
	"VPERMILPD",
	"VPERMILPS",
	"VPERMPD",
	"VPERMPS",
	"VPERMQ",
	"VPERMT2D",
	"VPERMT2PD",
	"VPERMT2PS",
	"VPERMT2Q",
	"VPERMT2W",
	"VPERMW",
	"VPEXPANDD",
	"VPEXPANDQ",
	"VPEXTRB",
	"VPEXTRD",
	"VPEXTRQ",
	"VPGATHERDD",
	"VPGATHERDQ",
	"VPGATHERQD",
	"VPGATHERQQ",
	"VPINSRD",
	"VPINSRQ",
	"VPLZCNTD",
	"VPLZCNTQ",
	"VPMADDUBSW",
	"VPMADDWD",
	"VPMAXSB",
	"VPMAXSD",
	"VPMAXSQ",
	"VPMAXSW",
	"VPMAXUB",
	"VPMAXUD",
	"VPMAXUQ",
	"VPMAXUW",
	"VPMINSB",
	"VPMINSD",
	"VPMINSQ",
	"VPMINSW",
	"VPMINUB",
	"VPMINUD",
	"VPMINUQ",
	"VPMINUW",
	"VPMOVB2M",
	"VPMOVD2M",
	"VPMOVDB",
	"VPMOVDW",
	"VPMOVM2B",
	"VPMOVM2D",
	"VPMOVM2Q",
	"VPMOVM2W",
	"VPMOVQ2M",
	"VPMOVQB",
	"VPMOVQD",
	"VPMOVQW",
	"VPMOVSDB",
	"VPMOVSDW",
	"VPMOVSQB",
	"VPMOVSQD",
	"VPMOVSQW",
	"VPMOVSWB",
	"VPMOVSXBD",
	"VPMOVSXBQ",
	"VPMOVSXBW",
	"VPMOVSXDQ",
	"VPMOVSXWD",
	"VPMOVSXWQ",
	"VPMOVUSDB",
	"VPMOVUSDW",
	"VPMOVUSQB",
	"VPMOVUSQD",
	"VPMOVUSQW",
	"VPMOVUSWB",
	"VPMOVW2M",
	"VPMOVWB",
	"VPMOVZXBD",
	"VPMOVZXBQ",
	"VPMOVZXBW",
	"VPMOVZXDQ",
	"VPMOVZXWD",
	"VPMOVZXWQ",
	"VPMULDQ",
	"VPMULHRSW",
	"VPMULHUW",
	"VPMULHW",
	"VPMULLD",
	"VPMULLQ",
	"VPMULLW",
	"VPMULUDQ",
	"VPORD",
	"VPORQ",
	"VPROLD",
	"VPROLQ",
	"VPROLVD",
	"VPROLVQ",
	"VPRORD",
	"VPRORQ",
	"VPRORVD",
	"VPRORVQ",
	"VPSADBW",
	"VPSCATTERDD",
	"VPSCATTERDQ",
	"VPSCATTERQD",
	"VPSCATTERQQ",
	"VPSHUFB",
	"VPSHUFD",
	"VPSHUFHW",
	"VPSHUFLW",
	"VPSLLD",
	"VPSLLDQ",
	"VPSLLQ",
	"VPSLLVD",
	"VPSLLVQ",
	"VPSLLVW",
	"VPSLLW",
	"VPSRAD",
	"VPSRAQ",
	"VPSRAVD",
	"VPSRAVQ",
	"VPSRAVW",
	"VPSRAW",
	"VPSRLD",
	"VPSRLDQ",
	"VPSRLQ",
	"VPSRLVD",
	"VPSRLVQ",
	"VPSRLVW",
	"VPSRLW",
	"VPSUBB",
	"VPSUBD",
	"VPSUBQ",
	"VPSUBSB",
	"VPSUBSW",
	"VPSUBUSB",
	"VPSUBUSW",
	"VPSUBW",
	"VPTERNLOGD",
	"VPTERNLOGQ",
	"VPTESTMB",
	"VPTESTMD",
	"VPTESTMQ",
	"VPTESTMW",
	"VPTESTNMB",
	"VPTESTNMD",
	"VPTESTNMQ",
	"VPTESTNMW",
	"VPUNPCKHBW",
	"VPUNPCKHDQ",
	"VPUNPCKHQDQ",
	"VPUNPCKHWD",
	"VPUNPCKLBW",
	"VPUNPCKLDQ",
	"VPUNPCKLQDQ",
	"VPUNPCKLWD",
	"VPXORD",
	"VPXORQ",
	"VRANGEPD",
	"VRANGEPS",
	"VRANGESD",
	"VRANGESS",
	"VRCP14PD",
	"VRCP14PS",
	"VRCP14SD",
	"VRCP14SS",
	"VREDUCEPD",
	"VREDUCEPS",
	"VREDUCESD",
	"VREDUCESS",
	"VRNDSCALEPD",
	"VRNDSCALEPS",
	"VRNDSCALESD",
	"VRNDSCALESS",
	"VRSQRT14PD",
	"VRSQRT14PS",
	"VRSQRT14SD",
	"VRSQRT14SS",
	"VSCALEFPD",
	"VSCALEFPS",
	"VSCALEFSD",
	"VSCALEFSS",
	"VSCATTERDPD",
	"VSCATTERDPS",
	"VSCATTERQPD",
	"VSCATTERQPS",
	"VSHUFF32X4",
	"VSHUFF64X2",
	"VSHUFI32X4",
	"VSHUFI64X2",
	"VSHUFPD",
	"VSHUFPS",
	"VSQRTPD",
	"VSQRTPS",
	"VSQRTSD",
	"VSQRTSS",
	"VSUBPD",
	"VSUBPS",
	"VSUBSD",
	"VSUBSS",
	"VUCOMISD",
	"VUCOMISS",
	"VUNPCKHPD",
	"VUNPCKHPS",
	"VUNPCKLPD",
	"VUNPCKLPS",
	"VXORPD",
	"VXORPS",
}

var v4Assembly = avx512Assembly

// VMOVNTDQ is problematic. it is both AVX, AVX512VL and AVX512F depending on what registers it operates on
// https://www.felixcloutier.com/x86/movntdq
var detectModeByRegisters = []string{
	"PUNPCKHQDQ",
	"VMOVDQU",
	"VMOVNTDQ",
	"VMOVNTDQA",
}

type AssemblyMode int8

const (
	na AssemblyMode = 0
	v1 AssemblyMode = 1
	v2 AssemblyMode = 2
	v3 AssemblyMode = 3
	v4 AssemblyMode = 4
)

var findRegisterExpression *regexp.Regexp
var findRegisterExpressionError error

func init() {
	findRegisterExpression, findRegisterExpressionError = regexp.Compile(`[^\(]*\(([^\(]*)\).*`)
	sort.Strings(x86SixyFourAssembly)
	sort.Strings(v2Assembly)
	sort.Strings(v3Assembly)
}

func pickRegisterName(operand string) string {
	if findRegisterExpressionError != nil {
		panic(findRegisterExpressionError)
	}

	result := findRegisterExpression.FindStringSubmatch(operand)
	if len(result) == 1 {
		return result[0]
	}

	return operand
}

// 	https://www.cryptologie.net/article/406/simd-instructions-in-go
var mmxRegisters = []string{
	"X0", "X1", "X2", "X3", "X4", "X5", "X6", "X7", "X8", "X9", "X10", "X11", "X12", "X13", "X14", "X15", "X16", "X17", "X18", "X19", "X20", "X21", "X22", "X23", "X24", "X25", "X26", "X28", "X29", "X30", "X31",
}
var avx512Registers = []string{
	"Z0", "Z1", "Z2", "Z3", "Z4", "Z5", "Z6", "Z7", "Z8", "Z9", "Z10", "Z11", "Z12", "Z13", "Z14", "Z15", "Z16", "Z17", "Z18", "Z19", "Z20", "Z21", "Z22", "Z23", "Z24", "Z25", "Z26", "Z28", "Z29", "Z30", "Z31",
}
var avxRegisters = []string{
	"Y0", "Y1", "Y2", "Y3", "Y4", "Y5", "Y6", "Y7", "Y8", "Y9", "Y10", "Y11", "Y12", "Y13", "Y14", "Y15", "Y16", "Y17", "Y18", "Y19", "Y20", "Y21", "Y22", "Y23", "Y24", "Y25", "Y26", "Y28", "Y29", "Y30", "Y31",
}

func Contains(collection []string, token string) bool {
	idx := sort.SearchStrings(collection, token)
	return idx != len(collection) && idx >= 0 && collection[idx] == token
}

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.BoolVar(&verbose, "verbose", false, "")

	var extended bool
	flag.BoolVar(&extended, "extended", false, "extended statistics")

	var printStatistics bool
	flag.BoolVar(&printStatistics, "s", false, "Print statistics")

	var inputFileName string
	flag.StringVar(&inputFileName, "input", "", "Input file name")
	flag.StringVar(&inputFileName, "i", "", "Input file name")

	flag.Parse()

	var globalMode AssemblyMode = na
	var operations = []int{0, 0, 0, 0}
	var v1Count map[string]int = make(map[string]int)
	var v2Count map[string]int = make(map[string]int)
	var v3Count map[string]int = make(map[string]int)
	var v4Count map[string]int = make(map[string]int)
	var scanner *bufio.Scanner
	if inputFileName != "" {
		reader, openErr := os.Open(inputFileName)
		defer reader.Close()
		if openErr != nil {
			log.Printf("Failed opening file %s\n", inputFileName)
			log.Panicln(openErr)
		}
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var context string = ""
	for scanner.Scan() {
		var mode AssemblyMode = na

		text := scanner.Text()
		if len(text) > 4 && text[:4] == "TEXT" {
			context = text[5:]
		} else {
			tokens := strings.Fields(text)
			var function string
			var instruction string
			for i, token := range tokens {
				if i == 0 {
					function = token
				}

				if Contains(x86SixyFourAssembly, token) && mode != v4 && mode != v3 && mode != v2 {
					mode = v1
					instruction = token
				}

				if Contains(v2Assembly, token) && mode != v4 && mode != v3 {
					mode = v2
					if verbose {
						fmt.Println("Found v2 instruction", token, "in function", function, context)
					}
					instruction = token

				}

				if Contains(v3Assembly, token) && mode != v4 {
					mode = v3
					if Contains(detectModeByRegisters, token) {
						var registerOneRaw = strings.TrimRight(tokens[i+1], ",")
						var registerTwoRaw = strings.TrimRight(tokens[i+2], ",")
						var registerOne = pickRegisterName(registerOneRaw)
						var registerTwo = pickRegisterName(registerTwoRaw)
						if Contains(avxRegisters, registerOne) || Contains(avxRegisters, registerTwo) {
							mode = v3
						} else {
							mode = v2
						}
					} else {
						mode = v3
					}

					if verbose {
						fmt.Println("Found v3 instruction", token, "in function", function, context)
					}
					instruction = token
				}

				if Contains(v4Assembly, token) {
					if Contains(detectModeByRegisters, token) {
						var registerOneRaw = strings.TrimRight(tokens[i+1], ",")
						var registerTwoRaw = strings.TrimRight(tokens[i+2], ",")
						var registerOne = pickRegisterName(registerOneRaw)
						var registerTwo = pickRegisterName(registerTwoRaw)
						if Contains(avx512Registers, registerOne) || Contains(avx512Registers, registerTwo) {
							mode = v4
						} else if Contains(avxRegisters, registerOne) || Contains(avxRegisters, registerTwo) {
							mode = v3
						} else {
							mode = v2
						}
					} else {
						mode = v4
					}

					if verbose {
						fmt.Println("Found v4 instruction", token, "in function", function, context)
					}
					instruction = token
				}
			}

			if mode == v1 {
				operations[0]++
				if val, ok := v1Count[instruction]; ok {
					v1Count[instruction] = val + 1
				} else {
					v1Count[instruction] = 1
				}
			} else if mode == v2 {
				operations[1]++
				if val, ok := v2Count[instruction]; ok {
					v2Count[instruction] = val + 1
				} else {
					v2Count[instruction] = 1
				}
			} else if mode == v3 {
				operations[2]++
				if val, ok := v3Count[instruction]; ok {
					v3Count[instruction] = val + 1
				} else {
					v3Count[instruction] = 1
				}
			} else if mode == v4 {
				operations[3]++
				if val, ok := v4Count[instruction]; ok {
					v4Count[instruction] = val + 1
				} else {
					v4Count[instruction] = 1
				}
			}
		}

		globalMode = AssemblyMode(math.Max(float64(mode), float64(globalMode)))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	if printStatistics {
		fmt.Println("x86", operations[0])
		if extended {
			for token, count := range v1Count {
				fmt.Println("    ", token, count)
			}
		}
		fmt.Println()

		fmt.Println("v2", operations[1])
		if extended {
			for token, count := range v2Count {
				fmt.Println("    ", token, count)
			}
		}
		fmt.Println()

		fmt.Println("v3", operations[2])
		if extended {
			for token, count := range v3Count {
				fmt.Println("    ", token, count)
			}
		}
		fmt.Println()

		fmt.Println("v4", operations[3])
		if extended {
			for token, count := range v4Count {
				fmt.Println("    ", token, count)
			}
		}
		fmt.Println()
	}

	if verbose {
		fmt.Printf("Minimum required GOAMD64=v%d\n", int(globalMode))
	} else {
		fmt.Printf("sGOAMD64=v%d\n", int(globalMode))
	}
}
