run this:
```
branch=$(git rev-parse --abbrev-ref HEAD) && \
upstream=$(git remote | head -n1) && \
git stash push -u -m "auto-stash before rebase" && \
git pull --rebase "$upstream" main && \
if git ls-remote --exit-code --heads "$upstream" "$branch" > /dev/null 2>&1; then \
  git push --force-with-lease; \
else \
  git push -u "$upstream" "$branch"; \
fi && \
git stash pop
```

Your commit message is "fix: ghissues"