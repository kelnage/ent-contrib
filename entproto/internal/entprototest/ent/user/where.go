// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// Unnecessary applies equality check predicate on the "unnecessary" field. It's identical to UnnecessaryEQ.
func Unnecessary(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUnnecessary, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// UnnecessaryEQ applies the EQ predicate on the "unnecessary" field.
func UnnecessaryEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUnnecessary, v))
}

// UnnecessaryNEQ applies the NEQ predicate on the "unnecessary" field.
func UnnecessaryNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUnnecessary, v))
}

// UnnecessaryIn applies the In predicate on the "unnecessary" field.
func UnnecessaryIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUnnecessary, vs...))
}

// UnnecessaryNotIn applies the NotIn predicate on the "unnecessary" field.
func UnnecessaryNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUnnecessary, vs...))
}

// UnnecessaryGT applies the GT predicate on the "unnecessary" field.
func UnnecessaryGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUnnecessary, v))
}

// UnnecessaryGTE applies the GTE predicate on the "unnecessary" field.
func UnnecessaryGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUnnecessary, v))
}

// UnnecessaryLT applies the LT predicate on the "unnecessary" field.
func UnnecessaryLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUnnecessary, v))
}

// UnnecessaryLTE applies the LTE predicate on the "unnecessary" field.
func UnnecessaryLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUnnecessary, v))
}

// UnnecessaryContains applies the Contains predicate on the "unnecessary" field.
func UnnecessaryContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUnnecessary, v))
}

// UnnecessaryHasPrefix applies the HasPrefix predicate on the "unnecessary" field.
func UnnecessaryHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUnnecessary, v))
}

// UnnecessaryHasSuffix applies the HasSuffix predicate on the "unnecessary" field.
func UnnecessaryHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUnnecessary, v))
}

// UnnecessaryIsNil applies the IsNil predicate on the "unnecessary" field.
func UnnecessaryIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUnnecessary))
}

// UnnecessaryNotNil applies the NotNil predicate on the "unnecessary" field.
func UnnecessaryNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUnnecessary))
}

// UnnecessaryEqualFold applies the EqualFold predicate on the "unnecessary" field.
func UnnecessaryEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUnnecessary, v))
}

// UnnecessaryContainsFold applies the ContainsFold predicate on the "unnecessary" field.
func UnnecessaryContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUnnecessary, v))
}

// HasBlogPosts applies the HasEdge predicate on the "blog_posts" edge.
func HasBlogPosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BlogPostsTable, BlogPostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlogPostsWith applies the HasEdge predicate on the "blog_posts" edge with a given conditions (other predicates).
func HasBlogPostsWith(preds ...predicate.BlogPost) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newBlogPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProfilePic applies the HasEdge predicate on the "profile_pic" edge.
func HasProfilePic() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProfilePicTable, ProfilePicColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfilePicWith applies the HasEdge predicate on the "profile_pic" edge with a given conditions (other predicates).
func HasProfilePicWith(preds ...predicate.Image) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newProfilePicStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSkipEdge applies the HasEdge predicate on the "skip_edge" edge.
func HasSkipEdge() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SkipEdgeTable, SkipEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkipEdgeWith applies the HasEdge predicate on the "skip_edge" edge with a given conditions (other predicates).
func HasSkipEdgeWith(preds ...predicate.SkipEdgeExample) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSkipEdgeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
