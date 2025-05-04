package datastar

import (
	"fmt"
)

// ValidFragmentMergeTypes is a list of valid fragment merge modes.
var ValidFragmentMergeTypes = []FragmentMergeMode{
	FragmentMergeModeMorph,
	FragmentMergeModeInner,
	FragmentMergeModeOuter,
	FragmentMergeModePrepend,
	FragmentMergeModeAppend,
	FragmentMergeModeBefore,
	FragmentMergeModeAfter,
	FragmentMergeModeUpsertAttributes,
}

// FragmentMergeTypeFromString converts a string to a [FragmentMergeMode].
func FragmentMergeTypeFromString(s string) (FragmentMergeMode, error) {
	switch s {
	case "morph":
		return FragmentMergeModeMorph, nil
	case "inner":
		return FragmentMergeModeInner, nil
	case "outer":
		return FragmentMergeModeOuter, nil
	case "prepend":
		return FragmentMergeModePrepend, nil
	case "append":
		return FragmentMergeModeAppend, nil
	case "before":
		return FragmentMergeModeBefore, nil
	case "after":
		return FragmentMergeModeAfter, nil
	case "upsertAttributes":
		return FragmentMergeModeUpsertAttributes, nil
	default:
		return "", fmt.Errorf("invalid fragment merge type: %s", s)
	}
}

// WithMergeMorph creates a MergeFragmentOption that merges fragments using the morph mode.
func WithMergeMorph() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModeMorph)
}

// WithMergeInner creates a MergeFragmentOption that merges fragments using the inner mode.
func WithMergeInner() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModeInner)
}

// WithMergeOuter creates a MergeFragmentOption that merges fragments using the outer mode.
func WithMergeOuter() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModeOuter)
}

// WithMergePrepend creates a MergeFragmentOption that merges fragments using the prepend mode.
func WithMergePrepend() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModePrepend)
}

// WithMergeAppend creates a MergeFragmentOption that merges fragments using the append mode.
func WithMergeAppend() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModeAppend)
}

// WithMergeBefore creates a MergeFragmentOption that merges fragments using the before mode.
func WithMergeBefore() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModeBefore)
}

// WithMergeAfter creates a MergeFragmentOption that merges fragments using the after mode.
func WithMergeAfter() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModeAfter)
}

// WithMergeUpsertAttributes creates a MergeFragmentOption that merges fragments using the upsert attributes mode.
func WithMergeUpsertAttributes() MergeFragmentOption {
	return WithMergeMode(FragmentMergeModeUpsertAttributes)
}

// WithSelectorID is a convenience wrapper for [WithSelector] option
// equivalent to calling `WithSelector("#"+id)`.
func WithSelectorID(id string) MergeFragmentOption {
	return WithSelector("#" + id)
}

// WithViewTransitions enables the use of view transitions when merging fragments.
func WithViewTransitions() MergeFragmentOption {
	return func(o *mergeFragmentOptions) {
		o.UseViewTransitions = true
	}
}

// WithoutViewTransitions disables the use of view transitions when merging fragments.
func WithoutViewTransitions() MergeFragmentOption {
	return func(o *mergeFragmentOptions) {
		o.UseViewTransitions = false
	}
}

// MergeFragmentf is a convenience wrapper for [MergeFragments] option
// equivalent to calling `MergeFragments(fmt.Sprintf(format, args...))`.
func (sse *ServerSentEventGenerator) MergeFragmentf(format string, args ...any) error {
	return sse.MergeFragments(fmt.Sprintf(format, args...))
}

// GetSSE is a convenience method for generating Datastar backend [get] action attribute.
//
// [get]: https://data-star.dev/reference/action_plugins#get
func GetSSE(urlFormat string, args ...any) string {
	return fmt.Sprintf(`@get('%s')`, fmt.Sprintf(urlFormat, args...))
}

// PostSSE is a convenience method for generating Datastar backend [post] action attribute.
//
// [post]: https://data-star.dev/reference/action_plugins#post
func PostSSE(urlFormat string, args ...any) string {
	return fmt.Sprintf(`@post('%s')`, fmt.Sprintf(urlFormat, args...))
}

// PutSSE is a convenience method for generating Datastar backend [put] action attribute.
//
// [put]: https://data-star.dev/reference/action_plugins#put
func PutSSE(urlFormat string, args ...any) string {
	return fmt.Sprintf(`@put('%s')`, fmt.Sprintf(urlFormat, args...))
}

// PatchSSE is a convenience method for generating Datastar backend [patch] action attribute.
//
// [patch]: https://data-star.dev/reference/action_plugins#patch
func PatchSSE(urlFormat string, args ...any) string {
	return fmt.Sprintf(`@patch('%s')`, fmt.Sprintf(urlFormat, args...))
}

// DeleteSSE is a convenience method for generating Datastar backend [delete] action attribute.
//
// [delete]: https://data-star.dev/reference/action_plugins#delete
func DeleteSSE(urlFormat string, args ...any) string {
	return fmt.Sprintf(`@delete('%s')`, fmt.Sprintf(urlFormat, args...))
}
