package ptx

import "fmt"

// SpecialReg represents PTX predefined special registers.
type SpecialReg int

const (
	// ---- Thread identification ----
	RegTidX SpecialReg = iota // %tid.x
	RegTidY                   // %tid.y
	RegTidZ                   // %tid.z

	// ---- CTA (block) dimensions ----
	RegNTidX // %ntid.x
	RegNTidY // %ntid.y
	RegNTidZ // %ntid.z

	// ---- Warp identification ----
	RegLaneId  // %laneid
	RegWarpId  // %warpid
	RegNWarpId // %nwarpid

	// ---- CTA (block) identification ----
	RegCTAIdX // %ctaid.x
	RegCTAIdY // %ctaid.y
	RegCTAIdZ // %ctaid.z

	// ---- Grid dimensions ----
	RegNCTAIdX // %nctaid.x
	RegNCTAIdY // %nctaid.y
	RegNCTAIdZ // %nctaid.z

	// ---- SM identification ----
	RegSMId  // %smid
	RegNSMId // %nsmid

	// ---- Grid identification ----
	RegGridId // %gridid

	// ---- Cluster identification (sm_90+) ----
	RegIsExplicitCluster // %is_explicit_cluster
	RegClusterIdX        // %clusterid.x
	RegClusterIdY        // %clusterid.y
	RegClusterIdZ        // %clusterid.z
	RegNClusterIdX       // %nclusterid.x
	RegNClusterIdY       // %nclusterid.y
	RegNClusterIdZ       // %nclusterid.z
	RegClusterCTAIdX     // %cluster_ctaid.x
	RegClusterCTAIdY     // %cluster_ctaid.y
	RegClusterCTAIdZ     // %cluster_ctaid.z
	RegClusterNCTAIdX    // %cluster_nctaid.x
	RegClusterNCTAIdY    // %cluster_nctaid.y
	RegClusterNCTAIdZ    // %cluster_nctaid.z
	RegClusterCTARank    // %cluster_ctarank
	RegClusterNCTARank   // %cluster_nctarank

	// ---- Lane masks ----
	RegLanemaskEq // %lanemask_eq
	RegLanemaskLe // %lanemask_le
	RegLanemaskLt // %lanemask_lt
	RegLanemaskGe // %lanemask_ge
	RegLanemaskGt // %lanemask_gt

	// ---- Clock ----
	RegClock   // %clock
	RegClockHi // %clock_hi
	RegClock64 // %clock64

	// ---- Performance monitoring (32-bit) ----
	RegPM0 // %pm0
	RegPM1 // %pm1
	RegPM2 // %pm2
	RegPM3 // %pm3
	RegPM4 // %pm4
	RegPM5 // %pm5
	RegPM6 // %pm6
	RegPM7 // %pm7

	// ---- Performance monitoring (64-bit) ----
	RegPM0_64 // %pm0_64
	RegPM1_64 // %pm1_64
	RegPM2_64 // %pm2_64
	RegPM3_64 // %pm3_64
	RegPM4_64 // %pm4_64
	RegPM5_64 // %pm5_64
	RegPM6_64 // %pm6_64
	RegPM7_64 // %pm7_64

	// ---- Shared memory ----
	RegDynamicSmemSize         // %dynamic_smem_size
	RegTotalSmemSize           // %total_smem_size
	RegAggrSmemSize            // %aggr_smem_size
	RegReservedSmemOffsetBegin // %reserved_smem_offset_begin
	RegReservedSmemOffsetEnd   // %reserved_smem_offset_end
	RegReservedSmemOffsetCap   // %reserved_smem_offset_cap
	RegReservedSmemOffset2     // %reserved_smem_offset_2

	// ---- Global timer ----
	RegGlobalTimer   // %globaltimer
	RegGlobalTimerLo // %globaltimer_lo
	RegGlobalTimerHi // %globaltimer_hi

	// ---- Execution graph ----
	RegCurrentGraphExec // %current_graph_exec

	// ---- Environment registers (%envreg0–%envreg31) ----
	// RegEnvReg0 is the base; specific indices are computed as RegEnvReg0 + n.
	RegEnvReg0
	RegEnvReg31 = RegEnvReg0 + 31
)

func (r SpecialReg) String() string {
	switch r {
	// Thread identification
	case RegTidX:
		return "%tid.x"
	case RegTidY:
		return "%tid.y"
	case RegTidZ:
		return "%tid.z"

	// CTA (block) dimensions
	case RegNTidX:
		return "%ntid.x"
	case RegNTidY:
		return "%ntid.y"
	case RegNTidZ:
		return "%ntid.z"

	// Warp identification
	case RegLaneId:
		return "%laneid"
	case RegWarpId:
		return "%warpid"
	case RegNWarpId:
		return "%nwarpid"

	// CTA (block) identification
	case RegCTAIdX:
		return "%ctaid.x"
	case RegCTAIdY:
		return "%ctaid.y"
	case RegCTAIdZ:
		return "%ctaid.z"

	// Grid dimensions
	case RegNCTAIdX:
		return "%nctaid.x"
	case RegNCTAIdY:
		return "%nctaid.y"
	case RegNCTAIdZ:
		return "%nctaid.z"

	// SM identification
	case RegSMId:
		return "%smid"
	case RegNSMId:
		return "%nsmid"

	// Grid identification
	case RegGridId:
		return "%gridid"

	// Cluster identification (sm_90+)
	case RegIsExplicitCluster:
		return "%is_explicit_cluster"
	case RegClusterIdX:
		return "%clusterid.x"
	case RegClusterIdY:
		return "%clusterid.y"
	case RegClusterIdZ:
		return "%clusterid.z"
	case RegNClusterIdX:
		return "%nclusterid.x"
	case RegNClusterIdY:
		return "%nclusterid.y"
	case RegNClusterIdZ:
		return "%nclusterid.z"
	case RegClusterCTAIdX:
		return "%cluster_ctaid.x"
	case RegClusterCTAIdY:
		return "%cluster_ctaid.y"
	case RegClusterCTAIdZ:
		return "%cluster_ctaid.z"
	case RegClusterNCTAIdX:
		return "%cluster_nctaid.x"
	case RegClusterNCTAIdY:
		return "%cluster_nctaid.y"
	case RegClusterNCTAIdZ:
		return "%cluster_nctaid.z"
	case RegClusterCTARank:
		return "%cluster_ctarank"
	case RegClusterNCTARank:
		return "%cluster_nctarank"

	// Lane masks
	case RegLanemaskEq:
		return "%lanemask_eq"
	case RegLanemaskLe:
		return "%lanemask_le"
	case RegLanemaskLt:
		return "%lanemask_lt"
	case RegLanemaskGe:
		return "%lanemask_ge"
	case RegLanemaskGt:
		return "%lanemask_gt"

	// Clock
	case RegClock:
		return "%clock"
	case RegClockHi:
		return "%clock_hi"
	case RegClock64:
		return "%clock64"

	// Performance monitoring (32-bit)
	case RegPM0:
		return "%pm0"
	case RegPM1:
		return "%pm1"
	case RegPM2:
		return "%pm2"
	case RegPM3:
		return "%pm3"
	case RegPM4:
		return "%pm4"
	case RegPM5:
		return "%pm5"
	case RegPM6:
		return "%pm6"
	case RegPM7:
		return "%pm7"

	// Performance monitoring (64-bit)
	case RegPM0_64:
		return "%pm0_64"
	case RegPM1_64:
		return "%pm1_64"
	case RegPM2_64:
		return "%pm2_64"
	case RegPM3_64:
		return "%pm3_64"
	case RegPM4_64:
		return "%pm4_64"
	case RegPM5_64:
		return "%pm5_64"
	case RegPM6_64:
		return "%pm6_64"
	case RegPM7_64:
		return "%pm7_64"

	// Shared memory
	case RegDynamicSmemSize:
		return "%dynamic_smem_size"
	case RegTotalSmemSize:
		return "%total_smem_size"
	case RegAggrSmemSize:
		return "%aggr_smem_size"
	case RegReservedSmemOffsetBegin:
		return "%reserved_smem_offset_begin"
	case RegReservedSmemOffsetEnd:
		return "%reserved_smem_offset_end"
	case RegReservedSmemOffsetCap:
		return "%reserved_smem_offset_cap"
	case RegReservedSmemOffset2:
		return "%reserved_smem_offset_2"

	// Global timer
	case RegGlobalTimer:
		return "%globaltimer"
	case RegGlobalTimerLo:
		return "%globaltimer_lo"
	case RegGlobalTimerHi:
		return "%globaltimer_hi"

	// Execution graph
	case RegCurrentGraphExec:
		return "%current_graph_exec"

	default:
		if r >= RegEnvReg0 && r <= RegEnvReg31 {
			return fmt.Sprintf("%%envreg%d", r-RegEnvReg0)
		}
		return "%unknown"
	}
}

// Type returns the PTX type of this special register.
func (r SpecialReg) Type() Type {
	switch r {
	case RegClock64, RegGlobalTimer, RegGlobalTimerLo, RegGlobalTimerHi, RegCurrentGraphExec:
		return U64
	case RegIsExplicitCluster:
		return Pred
	default:
		return U32
	}
}