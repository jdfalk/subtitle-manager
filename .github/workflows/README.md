# CI/CD Workflows

This project uses a modern CI/CD setup with separate workflows for frontend and
backend testing to ensure independent execution and clear separation of
concerns.

## Workflow Structure

### 1. Main CI Pipeline (`ci.yml`)

- **Purpose**: Orchestrates the entire CI process
- **Triggers**: Push to main, pull requests
- **Features**:
  - Detects changes to frontend/backend code
  - Runs only necessary workflows based on changes
  - Provides unified status reporting
  - Fails if any sub-workflow fails

### 2. Backend Workflow (`backend.yml`)

- **Purpose**: Tests Go backend code
- **Triggers**: Changes to `*.go`, `go.mod`, `go.sum`, `pkg/**`, `cmd/**`, etc.
- **Features**:
  - Go formatting check with `gofmt`
  - Static analysis with `go vet` and `staticcheck`
  - Unit tests with race detection and coverage
  - Database testing (PostgreSQL tests skip if unavailable)
  - Binary build verification
  - Coverage reporting to Codecov

### 3. Frontend Workflow (`frontend.yml`)

- **Purpose**: Tests Node.js frontend code
- **Triggers**: Changes to `webui/**`
- **Features**:
  - Multi-version Node.js testing (22, 24)
  - ESLint code quality checks
  - Prettier formatting verification
  - TypeScript type checking
  - Unit tests with Vitest
  - Production build verification
  - E2E tests with Playwright
  - Accessibility testing
  - Security audit with `npm audit`
  - Bundle size analysis

## Benefits

### Independent Execution

- Frontend and backend tests run separately
- Failure in one doesn't prevent the other from completing
- Faster feedback for developers working on specific areas

### Optimized Resource Usage

- Only runs necessary tests based on file changes
- Parallel execution where possible
- Efficient caching strategies

### Clear Status Reporting

- Individual status for frontend/backend
- Easy to identify which part of the codebase has issues
- Detailed artifacts and reports for debugging

## Local Development

### Backend Testing

```bash
# Run Go tests
go test ./...

# Run with coverage
go test -race -coverprofile=coverage.out ./...

# Format code
gofmt -s -w .

# Vet code
go vet ./...
```

### Frontend Testing

```bash
cd webui

# Install dependencies
npm install --legacy-peer-deps

# Run linting
npm run lint

# Check formatting
npm run format:check

# Fix formatting
npm run format

# Type checking
npm run type-check

# Run tests
npm test

# Run E2E tests
npm run test:e2e

# Build
npm run build
```

## Workflow Customization

### Adding New Checks

- **Backend**: Add steps to `backend.yml`
- **Frontend**: Add jobs to `frontend.yml`
- **Path Changes**: Update path filters in `ci.yml`

### Environment Variables

- Set secrets in GitHub repository settings
- Use `env:` in workflow files for configuration
- Database credentials are configured for PostgreSQL testing

### Artifacts

- Backend binary is uploaded as artifact
- Frontend build output is preserved
- Test reports and coverage data are stored
- Playwright reports for E2E test debugging

## Migration from Legacy

The previous `go.yml` workflow has been renamed to `go-legacy.yml` for
reference. The new system provides:

- ✅ Better separation of concerns
- ✅ Independent failure handling
- ✅ Optimized execution based on changes
- ✅ Enhanced frontend testing capabilities
- ✅ Comprehensive coverage reporting
- ✅ Modern tooling integration

## Troubleshooting

### Common Issues

1. **Dependency conflicts**: Use `npm install --legacy-peer-deps`
2. **Go module issues**: Run `go mod tidy`
3. **PostgreSQL connection**: Check service configuration in workflow
4. **Path filters**: Ensure correct file patterns in `ci.yml`

### Debugging

- Check individual workflow logs in GitHub Actions
- Download artifacts for detailed analysis
- Use `act` for local workflow testing
- Review coverage reports for test gaps
